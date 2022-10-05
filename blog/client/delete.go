package main

import (
	"context"
	"log"
	"time"

	pb "github.com/MelquesPaiva/grpc-course/blog/proto"
)

func deleteBlog(svc pb.BlogServiceClient, id string) {
	log.Println("deleteBlog was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &pb.BlogId{Id: id}

	_, err := svc.DeleteBlog(ctx, req)
	if err != nil {
		log.Fatalf("DeleteBlog failed: %v", err)
	}

	log.Println("DeleteBlog was successfull")
}
