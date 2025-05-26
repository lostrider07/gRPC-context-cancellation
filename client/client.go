package main

import (
	"context"
	"log"
	"time"

	pb "gRPC-context-cancellation/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)

	// Set a timeout shorter than server delay to trigger cancellation
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Alice"})
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	log.Printf("Response: %s", resp.Message)
}
