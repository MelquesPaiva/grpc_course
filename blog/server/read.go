package main

import (
	"context"

	pb "github.com/MelquesPaiva/grpc-course/blog/proto"
)

func (s *Server) ReadBlog(ctx context.Context, blog *pb.BlogId) (*pb.Blog, error) {
	return &pb.Blog{}, nil
}
