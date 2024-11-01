package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/vexxuh/magnetflow_filemanager/src/generated/src/protobuf"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedFileManagerServer
}

// Implement your service methods here

func main() {
	// Set up the gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterFileManagerServer(s, &server{})

	fmt.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
