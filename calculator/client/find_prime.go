package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Shashwat5522/go_gRPC/calculator/proto"
)

func findPrime(c pb.CalculatorServiceClient) {
	log.Println("find prime function called from client side")
	req := &pb.CalculatorRequest{
		Number: 120,
	}
	stream, err := c.FindPrime(context.Background(), req)
	if err != nil {
		log.Fatal("Error while calling find many times")
	}
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while  reading from stream:%v\n", err)
		}
		log.Printf("prime factor is :=%v\n",msg.Result)
	}
}
