package main

import (
	"context"
	"log"

	pb "github.com/Shashwat5522/go_gRPC/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("---delete blog was called from client side---")

	req := &pb.BlogId{
		Id: id,
	}

	_, err := c.DeleteBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("error occured while deleting blog:=%v\n", err)
	}
	log.Println("blog has deleted Sucessfully!!!")

}
