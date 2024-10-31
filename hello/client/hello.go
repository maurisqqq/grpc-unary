package main

import (
	"context"
	"log"

	pb "grpc-unary/hello/proto"
)

func doHello(c pb.HelloServiceClient) {
	log.Println("doHello was invoked")
	r, err := c.Hello(context.Background(), &pb.HelloRequest{FirstName: "Maulana"})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Hello %s\n", r.Result)
}
