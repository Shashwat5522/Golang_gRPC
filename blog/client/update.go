package main

import (
	"context"
	"log"

	pb "github.com/Shashwat5522/go_gRPC/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("**********Update Blog invoked************")
	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Clement",
		Title:    "A new Title",
		Content:  "Content of the first blog,with some awsome additions!!",
	}
	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error happened while updating:%v\n", err)

	}
	log.Println("blog was updated")
}
