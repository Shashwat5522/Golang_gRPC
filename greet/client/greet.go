package main

import(
	pb"github.com/Shashwat5522/go_gRPC/greet/proto"
	"log"
	"context"
)
 
func doGreet(c pb.GreetServiceClient){
	log.Println("doGreet was invoked")
	res,err:=c.Greet(context.Background(),&pb.GreetRequest{
		FirstName:"Shashwat",

	})
	if err!=nil{
		log.Fatalf("Could not greet:%v\n",err)

	}
	log.Printf("Greeting:%s\n",res.Result)
}