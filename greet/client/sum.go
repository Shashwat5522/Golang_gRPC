package main

import(
	pb"github.com/Shashwat5522/go_gRPC/greet/proto"
	"context"
	"log"

)

func doSum(c pb.GreetServiceClient){
	res,err:=c.Sum(context.Background(),&pb.SumRequest{
		Num1:3,
		Num2:4,

	})
	if err!=nil{
		log.Fatalf("could not sum:-%v",err)
	}
	log.Printf("sum is %v",res)
}