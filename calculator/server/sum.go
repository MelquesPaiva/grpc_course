package main

import (
	"context"
	"log"

	pb "github.com/MelquesPaiva/grpc-course/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoke with %v\n", req)

	return &pb.SumResponse{
		Result: req.FirstValue + req.SecondValue,
	}, nil
}
