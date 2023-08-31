package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Shashwat5522/go_gRPC/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "clement",
	}
	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatal("Error while calling GreetMany Times")
	}
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream:%v\n", err)
		}
		log.Printf("GreetMany times:%s\n",msg.Result)
	}
}
