package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

// Server - Unary Interceptor
func OrderUnarySeverInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	// Preprocessing logic
	// Get info about the current RPC call by examining the args passed in
	log.Println("====== [Server Interceptor] ", info.FullMethod)

	// Invoking the handler to complete the normal execution of a unary RPC.
	m, err := handler(ctx, req)

	// Post processing logic
	log.Printf(" Post Proc Message : %s", m)
	return m, err
}
