package main

import (
	"context"
	"fmt"
	"log"
	"math"

	pb "github.com/Shashwat5522/go_gRPC/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(context context.Context, request *pb.CalculatorRequest) (response *pb.CalculatorResponse, err error) {
	log.Println("Sqrt invoked with :%v\n", request)
	number := request.Number

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number:%d", number),
		)

	}
	return &pb.CalculatorResponse{
		Result: int32(math.Sqrt(float64(number))),
	}, nil
}
