package main

import (
	"io"
	"log"

	pb "github.com/Shashwat5522/go_gRPC/calculator/proto"
)

func (s *Server) FindMax(stream pb.CalculatorService_FindMaxServer) error {
	var maxi int32 = 0

	log.Println("find max function called from server side")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return  nil

		}
		if err != nil {
			log.Fatalf("Error while receiving :=%v\n", err)
		}
		if req.Number > maxi {
			maxi = req.Number
			err = stream.Send(&pb.CalculatorResponse{
				Result: maxi,
			})
			if err!=nil{
				log.Fatalf("error while sending response:=%v\n",err)
			}
		}


	}
}
