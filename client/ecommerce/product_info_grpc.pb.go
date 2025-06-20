// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: ecommerce/product_info.proto

package ecommerce

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ProductInfo_AddProduct_FullMethodName = "/ProductInfo/addProduct"
	ProductInfo_GetProduct_FullMethodName = "/ProductInfo/getProduct"
)

// ProductInfoClient is the client API for ProductInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductInfoClient interface {
	AddProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*ProductID, error)
	GetProduct(ctx context.Context, in *ProductID, opts ...grpc.CallOption) (*Product, error)
}

type productInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewProductInfoClient(cc grpc.ClientConnInterface) ProductInfoClient {
	return &productInfoClient{cc}
}

func (c *productInfoClient) AddProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*ProductID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductID)
	err := c.cc.Invoke(ctx, ProductInfo_AddProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productInfoClient) GetProduct(ctx context.Context, in *ProductID, opts ...grpc.CallOption) (*Product, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Product)
	err := c.cc.Invoke(ctx, ProductInfo_GetProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductInfoServer is the server API for ProductInfo service.
// All implementations must embed UnimplementedProductInfoServer
// for forward compatibility.
type ProductInfoServer interface {
	AddProduct(context.Context, *Product) (*ProductID, error)
	GetProduct(context.Context, *ProductID) (*Product, error)
	mustEmbedUnimplementedProductInfoServer()
}

// UnimplementedProductInfoServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProductInfoServer struct{}

func (UnimplementedProductInfoServer) AddProduct(context.Context, *Product) (*ProductID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProduct not implemented")
}
func (UnimplementedProductInfoServer) GetProduct(context.Context, *ProductID) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductInfoServer) mustEmbedUnimplementedProductInfoServer() {}
func (UnimplementedProductInfoServer) testEmbeddedByValue()                     {}

// UnsafeProductInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductInfoServer will
// result in compilation errors.
type UnsafeProductInfoServer interface {
	mustEmbedUnimplementedProductInfoServer()
}

func RegisterProductInfoServer(s grpc.ServiceRegistrar, srv ProductInfoServer) {
	// If the following call pancis, it indicates UnimplementedProductInfoServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProductInfo_ServiceDesc, srv)
}

func _ProductInfo_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductInfoServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductInfo_AddProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductInfoServer).AddProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductInfo_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductInfoServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductInfo_GetProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductInfoServer).GetProduct(ctx, req.(*ProductID))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductInfo_ServiceDesc is the grpc.ServiceDesc for ProductInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ProductInfo",
	HandlerType: (*ProductInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addProduct",
			Handler:    _ProductInfo_AddProduct_Handler,
		},
		{
			MethodName: "getProduct",
			Handler:    _ProductInfo_GetProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ecommerce/product_info.proto",
}

const (
	OrderManagement_GetOrder_FullMethodName      = "/OrderManagement/getOrder"
	OrderManagement_SearchOrders_FullMethodName  = "/OrderManagement/searchOrders"
	OrderManagement_UpdateOrders_FullMethodName  = "/OrderManagement/updateOrders"
	OrderManagement_ProcessOrders_FullMethodName = "/OrderManagement/processOrders"
)

// OrderManagementClient is the client API for OrderManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderManagementClient interface {
	GetOrder(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Order, error)
	SearchOrders(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Order], error)
	UpdateOrders(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[Order, wrapperspb.StringValue], error)
	ProcessOrders(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[wrapperspb.StringValue, CombinedShipment], error)
}

type orderManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderManagementClient(cc grpc.ClientConnInterface) OrderManagementClient {
	return &orderManagementClient{cc}
}

func (c *orderManagementClient) GetOrder(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Order, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Order)
	err := c.cc.Invoke(ctx, OrderManagement_GetOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderManagementClient) SearchOrders(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Order], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &OrderManagement_ServiceDesc.Streams[0], OrderManagement_SearchOrders_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[wrapperspb.StringValue, Order]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type OrderManagement_SearchOrdersClient = grpc.ServerStreamingClient[Order]

func (c *orderManagementClient) UpdateOrders(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[Order, wrapperspb.StringValue], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &OrderManagement_ServiceDesc.Streams[1], OrderManagement_UpdateOrders_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Order, wrapperspb.StringValue]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type OrderManagement_UpdateOrdersClient = grpc.ClientStreamingClient[Order, wrapperspb.StringValue]

func (c *orderManagementClient) ProcessOrders(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[wrapperspb.StringValue, CombinedShipment], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &OrderManagement_ServiceDesc.Streams[2], OrderManagement_ProcessOrders_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[wrapperspb.StringValue, CombinedShipment]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type OrderManagement_ProcessOrdersClient = grpc.BidiStreamingClient[wrapperspb.StringValue, CombinedShipment]

// OrderManagementServer is the server API for OrderManagement service.
// All implementations must embed UnimplementedOrderManagementServer
// for forward compatibility.
type OrderManagementServer interface {
	GetOrder(context.Context, *wrapperspb.StringValue) (*Order, error)
	SearchOrders(*wrapperspb.StringValue, grpc.ServerStreamingServer[Order]) error
	UpdateOrders(grpc.ClientStreamingServer[Order, wrapperspb.StringValue]) error
	ProcessOrders(grpc.BidiStreamingServer[wrapperspb.StringValue, CombinedShipment]) error
	mustEmbedUnimplementedOrderManagementServer()
}

// UnimplementedOrderManagementServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOrderManagementServer struct{}

func (UnimplementedOrderManagementServer) GetOrder(context.Context, *wrapperspb.StringValue) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOrderManagementServer) SearchOrders(*wrapperspb.StringValue, grpc.ServerStreamingServer[Order]) error {
	return status.Errorf(codes.Unimplemented, "method SearchOrders not implemented")
}
func (UnimplementedOrderManagementServer) UpdateOrders(grpc.ClientStreamingServer[Order, wrapperspb.StringValue]) error {
	return status.Errorf(codes.Unimplemented, "method UpdateOrders not implemented")
}
func (UnimplementedOrderManagementServer) ProcessOrders(grpc.BidiStreamingServer[wrapperspb.StringValue, CombinedShipment]) error {
	return status.Errorf(codes.Unimplemented, "method ProcessOrders not implemented")
}
func (UnimplementedOrderManagementServer) mustEmbedUnimplementedOrderManagementServer() {}
func (UnimplementedOrderManagementServer) testEmbeddedByValue()                         {}

// UnsafeOrderManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderManagementServer will
// result in compilation errors.
type UnsafeOrderManagementServer interface {
	mustEmbedUnimplementedOrderManagementServer()
}

func RegisterOrderManagementServer(s grpc.ServiceRegistrar, srv OrderManagementServer) {
	// If the following call pancis, it indicates UnimplementedOrderManagementServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OrderManagement_ServiceDesc, srv)
}

func _OrderManagement_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderManagement_GetOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServer).GetOrder(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderManagement_SearchOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(wrapperspb.StringValue)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderManagementServer).SearchOrders(m, &grpc.GenericServerStream[wrapperspb.StringValue, Order]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type OrderManagement_SearchOrdersServer = grpc.ServerStreamingServer[Order]

func _OrderManagement_UpdateOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderManagementServer).UpdateOrders(&grpc.GenericServerStream[Order, wrapperspb.StringValue]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type OrderManagement_UpdateOrdersServer = grpc.ClientStreamingServer[Order, wrapperspb.StringValue]

func _OrderManagement_ProcessOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderManagementServer).ProcessOrders(&grpc.GenericServerStream[wrapperspb.StringValue, CombinedShipment]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type OrderManagement_ProcessOrdersServer = grpc.BidiStreamingServer[wrapperspb.StringValue, CombinedShipment]

// OrderManagement_ServiceDesc is the grpc.ServiceDesc for OrderManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OrderManagement",
	HandlerType: (*OrderManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getOrder",
			Handler:    _OrderManagement_GetOrder_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "searchOrders",
			Handler:       _OrderManagement_SearchOrders_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "updateOrders",
			Handler:       _OrderManagement_UpdateOrders_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "processOrders",
			Handler:       _OrderManagement_ProcessOrders_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ecommerce/product_info.proto",
}
