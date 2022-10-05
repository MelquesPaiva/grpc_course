package main

import (
	"context"
	"log"
	"time"

	pb "github.com/MelquesPaiva/grpc-course/blog/proto"
	"google.golang.org/grpc/status"
)

func readBlog(svc pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("readBlog was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &pb.BlogId{Id: id}
	res, err := svc.ReadBlog(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Fatalf("Unexpected gRPC error: %v\n", e)
		}
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog was read: %v", res)
	return res
}
