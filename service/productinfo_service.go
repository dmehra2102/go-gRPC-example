package main

import (
	"context"
	"fmt"
	"log"
	pb "productInfo/service/ecommerce"
	"strings"

	"github.com/gofrs/uuid/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type server struct {
	pb.UnimplementedProductInfoServer
	pb.UnimplementedOrderManagementServer
	productMap map[string]*pb.Product
	orderMap map[string]*pb.Order
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
func (s *server) GetOrder(ctx context.Context, orderId *wrapperspb.StringValue) (*pb.Order, error){
	ord := s.orderMap[orderId.Value]
	return ord, nil
}

func (s *server) SearchOrders(searchQuery *wrapperspb.StringValue, stream pb.OrderManagement_SearchOrdersServer) error {
	for key,order := range s.orderMap {
		log.Print(key, order)
		for _, itemStr := range order.Items {
			log.Print(itemStr)
			if strings.Contains(itemStr, searchQuery.Value) {
				err := stream.Send(order)
				if err != nil {
				   return fmt.Errorf("error sending message to stream : %v",err) 
				}
				log.Print("Matching Order Found : " + key)
				break
			}
		}
	}
	return nil
}