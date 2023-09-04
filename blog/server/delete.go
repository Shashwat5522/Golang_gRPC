package main

import (
	"context"
	"log"

	pb "github.com/Shashwat5522/go_gRPC/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {
	log.Println("delete blog invoked from server side")
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"cannot parse ID",
		)
	}
	
	filter := bson.M{"_id": oid}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"cannot delete blog with this id",
		)
	}

	log.Printf("%v is delete sucessfully!!!", res)

	return &emptypb.Empty{}, nil
}
