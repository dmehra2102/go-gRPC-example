package main

import (
	"log"
	"time"

	"google.golang.org/grpc"
)

type wrappedStream struct {
	grpc.ServerStream
}

func (w *wrappedStream) SendMsg(m any) error {
	log.Printf("===== [Server Stream Interceptor Wrapper] "+"Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.SendMsg(m)
}

func (w *wrappedStream) RecvMsg(m any) error {
	log.Printf("===== [Server Stream Interceptor Wrapper] "+"Recieve a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.RecvMsg(m)
}

func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

// StreamServerInterceptor func(srv any, ss ServerStream, info *StreamServerInfo, handler StreamHandler) error

func OrderStreamServerInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("====== [Server Stream Interceptor] ", info.FullMethod)
	err := handler(srv, newWrappedStream(ss))
	if err != nil {
		log.Printf("RPC failed with error %v", err)
	}
	return err
}
