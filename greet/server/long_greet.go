package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/Shashwat5522/go_gRPC/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {

	log.Println("LongGreet function was called")
	res := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatal("Error while reading client stream")
		}
		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}
