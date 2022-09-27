package main

import (
	"context"

	pb "github.com/MelquesPaiva/grpc-course/blog/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *Server) UpdateBlog(ctx context.Context, blogId *pb.BlogId) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
