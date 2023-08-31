package main

import (
	"context"
	"log"

	pb "github.com/Shashwat5522/go_gRPC/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("dosqrt was invoked")
	res, err := c.Sqrt(context.Background(), &pb.CalculatorRequest{
		Number: n,
	})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error message from Server:%s\n", e.Message())
			log.Printf("Error code from server:%s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Println("we probably sent a negative number")
				return
			}

		} else {
			log.Fatalf("A non gRPC error:%v\n")
		}

	}
	log.Printf("Sqrt:%v\n", res.Result)
}
