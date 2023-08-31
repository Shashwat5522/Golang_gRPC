package main

import (
	"io"
	"log"

	pb "github.com/Shashwat5522/go_gRPC/greet/proto"
)

func (s *Server) BidirectionalGreet(stream pb.GreetService_BidirectionalGreetServer) error {
	log.Println("Bidirectional Greet was invoked")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatal("Error while reading client stream")

		}
		res := "Hello" + req.FirstName + "!"
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})
		if err != nil {
			log.Fatalf("Error while sending data to client:%v\n", err)
		}

	}
}
