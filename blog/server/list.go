package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Shashwat5522/go_gRPC/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListBlog(in *emptypb.Empty, stream pb.BlogService_ListBlogServer) error {
	log.Println("ListBlog was invoked")

	curr, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error:%v", err),
		)
	}

	defer curr.Close(context.Background())

	for curr.Next(context.Background()) {
		data := &BlogItem{}
		err := curr.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MongoDB %v",err),
			)
		
		}
		stream.Send(documentToBlog(data))

	}
	if err=curr.Err();err!=nil{
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error:= %v",err),


		)
	}
	return nil
}
