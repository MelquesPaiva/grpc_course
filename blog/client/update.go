package main

import (
	"context"
	"log"
	"time"

	pb "github.com/MelquesPaiva/grpc-course/blog/proto"
)

func updateBlog(svc pb.BlogServiceClient, id string) {
	log.Println("Update blog was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Melques Paiva",
		Title:    "Update Title",
		Content:  "Update with additions",
	}

	_, err := svc.UpdateBlog(ctx, newBlog)
	if err != nil {
		log.Fatalf("Error while updating a blog: %v", err)
	}

	log.Println("Blog was updated")
}
