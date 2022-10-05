package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/MelquesPaiva/grpc-course/blog/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListBlogs(in *empty.Empty, stream pb.BlogService_ListBlogsServer) error {
	log.Printf("ListBlogs was invoked")
	ctx := context.Background()

	res, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v", err))
	}
	defer res.Close(ctx)

	for res.Next(ctx) {
		var current BlogItem
		if err = res.Decode(&current); err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Internal error: %v", err),
			)
		}

		stream.Send(documentToBlog(current))
	}

	if err := res.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}

	return nil
}
