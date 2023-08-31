package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Shashwat5522/go_gRPC/greet/proto"
)

func bidireactionalGreet(c pb.GreetServiceClient) {
	log.Println("bidirectional greet was invoked from client")

	stream, err := c.BidirectionalGreet(context.Background())
	if err != nil {
		log.Fatalf("error while creating stream:%v\n", err)

	}
	reqs := []*pb.GreetRequest{
		{FirstName: "Harry"},
		{FirstName: "Clement"},
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
				// log.Printf("Error while receiving:%v\n", err)
				break
			}
			if err!=nil{
				log.Printf("Error while receiving:%v\n",err)
				break
			}
			log.Printf("Received:%v\n", res.Result)
			
		}
		close(waitc)
	}()
	<-waitc
	
}
