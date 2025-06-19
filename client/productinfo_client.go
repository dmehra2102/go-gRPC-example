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
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
}
