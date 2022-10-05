package main

import (
	"context"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/MelquesPaiva/grpc-course/blog/proto"
	"github.com/MelquesPaiva/grpc-course/blog/server/db"
	"github.com/golang/protobuf/ptypes/empty"
)

var collection *mongo.Collection

const (
	port = ":50053"
)

type CrudInterface interface {
	CreateBlog(ctx context.Context, blog *pb.Blog) (*pb.BlogId, error)
	DeleteBlog(ctx context.Context, blogId *pb.BlogId) (*empty.Empty, error)
	ListBlogs(in *empty.Empty, stream pb.BlogService_ListBlogsServer) error
	UpdateBlog(ctx context.Context, blog *pb.Blog) (*empty.Empty, error)
	ReadBlog(ctx context.Context, blog *pb.BlogId) (*pb.Blog, error)
}

type Server struct {
	pb.BlogServiceServer
	db *mongo.Client
}

func main() {
	ctx := context.Background()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening on %v: %v", port, err)
	}

	log.Printf("Listening on %s\n", port)

	dbInstance := db.NewDb()
	client, err := dbInstance.Connect()
	if err != nil {
		log.Fatalf("Error connecting to the client database: %v\n", err)
	}

	s := &Server{}
	s.db = client

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalf("Error disconnecting to the client database: %v\n", err)
		}
	}()

	// Caso não exista o database blogdb, ele será criado
	// Caso não exista a collection blog, ela será criada
	collection = client.Database("blogdb").Collection("blog")

	srv := grpc.NewServer()
	pb.RegisterBlogServiceServer(srv, s)
	reflection.Register(srv)

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Fail to serve: %v\n", err)
	}
}
