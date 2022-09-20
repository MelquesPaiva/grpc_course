package main

import (
	"log"
	"net"

	pb "github.com/MelquesPaiva/grpc-course/greet/proto"
	"google.golang.org/grpc"
)

var s *grpc.Server

const (
	port = ":50051"
)

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Fail to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", port)

	s = grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Fail to serve: %v\n", err)
	}
}
