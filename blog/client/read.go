package main

import (
	"log"
	"context"
	pb "github.com/Shashwat5522/go_gRPC/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("-------readBlog was invoked------")
	req:=&pb.BlogId{
		Id:id,
	}

	res,err:=c.ReadBlog(context.Background(),req)

	if err!=nil{
		log.Fatalf("Error happened while reading:-%v\n",err)
	}
	log.Printf("Blog was read:%v\n",res)
	return res
}
