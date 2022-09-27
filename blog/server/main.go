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
)

var collection *mongo.Collection

const (
	port = ":50053"
)

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
