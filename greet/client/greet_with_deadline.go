package main

import (
	"log"
	"time"

	"context"

	pb "github.com/Shashwat5522/go_gRPC/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetwith Deadline was invoked at client side")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Clement",
	}
	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {		
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded!!")
				return
			} else {
				log.Fatalf("Unexpected gRPC error:=%v\n", err)
			}
		} else {
			log.Fatal("A non gRPC error:%v\n", err)
		}
	}

	log.Printf("GreetithDeadline:%s\n", res.Result)

}
