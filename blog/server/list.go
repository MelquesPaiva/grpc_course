package main

import (
	pb "github.com/MelquesPaiva/grpc-course/blog/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *Server) ListBlogs(in *empty.Empty, stream pb.BlogService_ListBlogsServer) error {
	return nil
}
