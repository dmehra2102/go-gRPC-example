package main

import (
	"context"
	"fmt"
	"io"
	"log"
	pb "productInfo/service/ecommerce"
	"strconv"
	"strings"

	"github.com/gofrs/uuid/v5"
	erp "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type server struct {
	pb.UnimplementedProductInfoServer
	pb.UnimplementedOrderManagementServer
	productMap map[string]*pb.Product
	orderMap   map[string]*pb.Order
}

// AddProduct implements ecommerce.AddProduct
func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal,
			"Error while generating Product ID : %v", err)
	}
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[in.Id] = in
	return &pb.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
}

// GetProduct implements ecommerce.GetProduct
func (s *server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	value, exists := s.productMap[in.Value]
	if exists {
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Product does not exist with ID : %s.", in.Value)
}

// GetOrder implements ecommerce.GetOrder
func (s *server) GetOrder(ctx context.Context, orderId *wrapperspb.StringValue) (*pb.Order, error) {
	i, err := strconv.ParseInt(orderId.Value, 10, 64)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid order ID: %v", err)
	}
	if i < 0 {
		log.Printf("Order ID is invalid! -> Received Order ID %s", orderId.Value)

		errorStatus := status.New(codes.InvalidArgument, "Invalid information received")
		ds, err := errorStatus.WithDetails(&erp.BadRequest_FieldViolation{
			Field: "ID",
			Description: fmt.Sprintf(
				"Order ID received is not valid %s ", orderId.Value),
		})
		if err != nil {
			return nil, errorStatus.Err()
		}

		return nil, ds.Err()
	}
	ord := s.orderMap[orderId.Value]
	return ord, nil
}

// SearchOrders implements ecommerce.SearchOrders
func (s *server) SearchOrders(searchQuery *wrapperspb.StringValue, stream pb.OrderManagement_SearchOrdersServer) error {
	for key, order := range s.orderMap {
		log.Print(key, order)
		for _, itemStr := range order.Items {
			log.Print(itemStr)
			if strings.Contains(itemStr, searchQuery.Value) {
				err := stream.Send(order)
				if err != nil {
					return fmt.Errorf("error sending message to stream : %v", err)
				}
				log.Print("Matching Order Found : " + key)
				break
			}
		}
	}
	return nil
}

// UpdateOrders implements ecommerce.UpdateOrders
func (s *server) UpdateOrders(stream pb.OrderManagement_UpdateOrdersServer) error {
	ordersStr := "updated Order IDs : "
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&wrapperspb.StringValue{Value: "Orders processed" + ordersStr})
		}

		s.orderMap[order.Id] = order
		log.Printf("Order ID : %s : Updated", order.Id)
		ordersStr += order.Id + ", "
	}
}
