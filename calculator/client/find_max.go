package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Shashwat5522/go_gRPC/calculator/proto"
)

func findMax(c pb.CalculatorServiceClient) {
	log.Println("find Max function invoked from client side")
	stream, err := c.FindMax(context.Background())
	if err != nil {
		log.Fatalf("error while receiving :=%v\n", err)
	}
	reqs := []*pb.CalculatorRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
	}
	waitc := make(chan struct{})
	go func() {
		for _, req := range reqs {
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()

	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error is:=%v\n", err)
				break
			}
			log.Printf("max number is:=%v\n", res.Result)

		}
		close(waitc)
	}()
	<-waitc

}
