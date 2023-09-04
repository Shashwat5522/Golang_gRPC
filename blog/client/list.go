package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Shashwat5522/go_gRPC/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlogs(c pb.BlogServiceClient) {
	log.Println("####list blog function called from client side####")

	stream, err := c.ListBlog(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error while calling listblog:%v\n", err)

	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("something happened:%v\n", err)
		}
		log.Println(res)

	}

}
