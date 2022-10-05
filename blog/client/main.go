package main

import (
	"log"

	pb "github.com/MelquesPaiva/grpc-course/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	addr = "localhost:50053"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Fail to connect: %v\n", err)
	}
	defer conn.Close()

	svc := pb.NewBlogServiceClient(conn)
	id := createBlog(svc)
	readBlog(svc, id)
	updateBlog(svc, id)
	deleteBlog(svc, id)
	listBlogs(svc)
}
