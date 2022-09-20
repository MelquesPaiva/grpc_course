package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/MelquesPaiva/grpc-course/calculator/proto"
)

var s *grpc.Server

const (
	port = ":50052"
)

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening on %v: %v", port, err)
	}

	log.Printf("Listening on %s\n", port)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Fail to serve: %v\n", err)
	}
}
