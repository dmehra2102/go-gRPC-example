package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func OrderUnaryClientInterceptor(ctx context.Context, method string, req, reply any,
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// Preprocessor phase
	log.Println("Method : " + method)

	// Invoking the remote method
	err := invoker(ctx, method, req, reply, cc, opts...)

	// Postprocessor phase
	log.Println(reply)

	return err
}
