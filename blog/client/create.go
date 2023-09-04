package main

import (
	"log"
	"context"
	pb "github.com/Shashwat5522/go_gRPC/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("create blog was invoked")

	blog:=&pb.Blog{
		AuthorId:"Clement",  
		Title:"My first blog",
		Content:"content of the first blog",

	}
	res,err:=c.CreateBlog(context.Background(),blog)
	if err!=nil{
		log.Fatalf("unexpected error:%v\n",err)

	}
	log.Printf("Blog has been created:%s\n",res.Id)
	return res.Id
}
