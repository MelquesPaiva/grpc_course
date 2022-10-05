package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/MelquesPaiva/grpc-course/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlogs(svc pb.BlogServiceClient) {
	log.Printf("listBlogs was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := svc.ListBlogs(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	for {
		blog, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Unexpected error reading the stream: %v\n", err)
		}

		log.Println(blog)
	}
}
