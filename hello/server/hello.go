package main

import (
	"context"
	pb "grpc-unary/hello/proto"
	"log"
)

func (s *Server) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Greet was invoked with %v\n", in)
	return &pb.HelloResponse{Result: "Hello " + in.FirstName}, nil
}
