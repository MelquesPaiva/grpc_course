package main

import (
	"context"
	"log"

	pb "github.com/MelquesPaiva/grpc-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UpdateBlog(ctx context.Context, blog *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("Update blog was invoked: %v", blog)

	id, err := primitive.ObjectIDFromHex(blog.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse blog ID",
		)
	}

	data := &BlogItem{
		AuthorId: blog.AuthorId,
		Title:    blog.Title,
		Content:  blog.Content,
	}

	res, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": data},
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not update")
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "The blog was not found")
	}

	return &emptypb.Empty{}, nil
}
