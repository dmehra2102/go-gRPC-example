package main

import (
	"context"
	"io"
	"log"
	pb "productinfo/client/ecommerce"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.NewClient(address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(OrderUnaryClientInterceptor),
		grpc.WithStreamInterceptor(OrderStreamClientInterceptor),
	)

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewProductInfoClient(conn)
	orderMgtClient := pb.NewOrderManagementClient(conn)

	name := "Apple iPhone 11"
	description := `Meet Apple iPhone 11. All-new dual-camera system with
              Ultra Wide and Night mode.`
	price := float32(1000.0)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// ADD Product
	r, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description, Price: price})

	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	log.Printf("Product ID: %s added successfully", r.Value)

	// GET Product
	product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	if err != nil {
		log.Fatalf("Could not get product: %v", err)
	}
	log.Printf("Product: %v", product.String())

	// GET order
	retrievedOrder, err := orderMgtClient.GetOrder(ctx,
		&wrapperspb.StringValue{Value: "106"})
	if err != nil {
		log.Fatalf("Could not get product: %v", err)
	}
	log.Print("GetOrder Response -> ", retrievedOrder)

	// Search Orders
	searchStream, _ := orderMgtClient.SearchOrders(ctx, &wrapperspb.StringValue{Value: "Google"})
	for {
		searchOrder, err := searchStream.Recv()
		if err == io.EOF {
			break
		}

		log.Print("Search Result : ", searchOrder)
	}

	// Define updated orders
	updOrder1 := pb.Order{
		Id:          "106",
		Items:       []string{"Apple iPhone 11", "Apple Watch"},
		Destination: "San Jose, CA",
		Price:       1200.0,
	}
	updOrder2 := pb.Order{
		Id:          "107",
		Items:       []string{"Google Pixel 4A"},
		Destination: "Mountain View, CA",
		Price:       800.0,
	}
	updOrder3 := pb.Order{
		Id:          "108",
		Items:       []string{"OnePlus 8"},
		Destination: "Sunnyvale, CA",
		Price:       700.0,
	}

	// Update Orders
	updateStream, err := orderMgtClient.UpdateOrders(ctx)
	if err != nil {
		log.Fatalf("%v.UpdateOrders(_) = _, %v", orderMgtClient, err)
	}

	// Updating order 1
	if err := updateStream.Send(&updOrder1); err != nil {
		log.Fatalf("%v.Send(OrderID=%v) = %v", updateStream, updOrder1.Id, err)
	}

	// Updating order 2
	if err := updateStream.Send(&updOrder2); err != nil {
		log.Fatalf("%v.Send(OrderID=%v) = %v", updateStream, updOrder2.Id, err)
	}

	// Updating order 3
	if err := updateStream.Send(&updOrder3); err != nil {
		log.Fatalf("%v.Send(OrderID=%v) = %v", updateStream, updOrder3.Id, err)
	}

	updateRes, err := updateStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", updateStream, err, nil)
	}
	log.Printf("Update Orders Res : %s", updateRes)
}
