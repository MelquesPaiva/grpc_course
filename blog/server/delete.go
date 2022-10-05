package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/MelquesPaiva/grpc-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, blogId *pb.BlogId) (*emptypb.Empty, error) {
	log.Printf("DeleteBlog was invoked with: %v\n", blogId)

	var deleted bson.M
	emp := &emptypb.Empty{}

	id, err := primitive.ObjectIDFromHex(blogId.Id)
	if err != nil {
		return emp, status.Errorf(codes.Internal, "Could not parse ID")
	}

	err = collection.FindOneAndDelete(ctx, bson.M{
		"_id": id,
	}).Decode(&deleted)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return emp, status.Errorf(
				codes.NotFound,
				fmt.Sprintf("The blog was not found: %v", err),
			)
		}

		return emp, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v\n", err),
		)
	}

	return emp, nil
}
