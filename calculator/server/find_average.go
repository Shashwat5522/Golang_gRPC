package main

import (
	"io"
	"log"

	pb "github.com/Shashwat5522/go_gRPC/calculator/proto"
)

func (s *Server) FindAverage(stream pb.CalculatorService_FindAverageServer) error {
	log.Println("Find average function called from server!!")
	var res int32= 0
	var i int32= 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.CalculatorResponse{
				Result: res/i,
			})
		}
		if err != nil {
			log.Fatal("Error while reading client stream")
		}
		res += req.Number
		i++

	}

}
