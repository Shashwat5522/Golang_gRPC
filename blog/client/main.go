package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Shashwat5522/go_gRPC/blog/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect:%v\n", err)

	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)
	// id := createBlog(c)
	// readBlog(c, id)
	// // readBlog(c,"aNonExistingIDs")
	// updateBlog(c, id)
	// listBlogs(c)
	deleteBlog(c,"64f571521e4917362c2bcb56")

}
