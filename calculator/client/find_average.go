package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Shashwat5522/go_gRPC/calculator/proto"
)

func findAvg(c pb.CalculatorServiceClient) {
	log.Println("findAvg function invoked from client side")

	reqs := []*pb.CalculatorRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
	}
	stream, err := c.FindAverage(context.Background())
	if err != nil {
		log.Fatal("Error while calling FindAverage %v\n", err)
	}
	for _, req := range reqs {
		log.Printf("Sending req:%v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from server:%v\n", err)

	}
	log.Printf("average is :%v\n", resp.Result)
}
