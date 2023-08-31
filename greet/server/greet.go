package main

import(
	"context"
	pb "github.com/Shashwat5522/go_gRPC/greet/proto"
	"log"
)

func (s *Server) Greet(context context.Context, request *pb.GreetRequest) (*pb.GreetResponse,error){
	log.Printf("Greet function was invoked with %v\n",request)
	return &pb.GreetResponse{
		Result:"Hello "+request.FirstName,
	},nil
}
