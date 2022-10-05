package main

import (
	"context"
	"log"
	"time"

	pb "github.com/MelquesPaiva/grpc-course/blog/proto"
	"google.golang.org/grpc/status"
)

func createBlog(svc pb.BlogServiceClient) string {
	log.Printf("createBlog was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := svc.CreateBlog(ctx, &pb.Blog{
		AuthorId: "Melques Santos",
		Title:    "GRPC-COURSE Client test2",
		Content:  "Testando conteúdo da nova solicilitação do blog2",
	})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Fatalf("Unexpected gRPC error: %v\n", e)
		}
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog created with success status: Id = %v\n", res.Id)
	return res.Id
}
