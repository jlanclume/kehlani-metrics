package main

import (
	"api/proto"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedMetricEventServiceServer
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Enable reflection for testing
	reflection.Register(grpcServer)

	// Register the metric service with the server
	proto.RegisterMetricEventServiceServer(grpcServer, &server{})

	fmt.Println("Starting server on port 50051...")
	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
