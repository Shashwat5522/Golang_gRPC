package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Shashwat5522/go_gRPC/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, request *pb.GreetRequest) (*pb.GreetResponse, error) {

	log.Printf("GreetWithDeadline was invoked with %v\n", request)
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("the client canceled the request")
			return nil, status.Error(codes.Canceled, "the client canceled the request")

		}
		time.Sleep(1 * time.Second)
	}
	return &pb.GreetResponse{
		Result: "Hello" + request.FirstName,
	}, nil
}
  