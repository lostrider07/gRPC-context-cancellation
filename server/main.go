package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "gRPC-context-cancellation/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received request for: %s", req.Name)

	// Simulate long work with cancellation awareness
	select {
	case <-time.After(5 * time.Second): // simulate slow processing
		return &pb.HelloResponse{
			Message: fmt.Sprintf("Hello, %s!", req.Name),
		}, nil
	case <-ctx.Done():
		log.Printf("Request cancelled: %v", ctx.Err())
		return nil, ctx.Err()
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	log.Println("gRPC server running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
