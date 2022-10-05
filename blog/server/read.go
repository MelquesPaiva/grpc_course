package main

import (
	"context"
	"fmt"

	pb "github.com/MelquesPaiva/grpc-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadBlog(ctx context.Context, blog *pb.BlogId) (*pb.Blog, error) {
	result := BlogItem{}

	id, err := primitive.ObjectIDFromHex(blog.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse blog ID",
		)
	}

	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(
				codes.NotFound,
				fmt.Sprintf("The blog was not found: %v", err),
			)
		}

		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	return documentToBlog(result), nil
}
