package main

import(
	"context"
	pb"github.com/Shashwat5522/go_gRPC/greet/proto"
)

func(s *Server) Sum(context context.Context,request *pb.SumRequest)(*pb.SumResponse,error){
	num1:=request.Num1
	num2:=request.Num2
	
	return &pb.SumResponse{
		Sum:num1+num2,
	},nil

}