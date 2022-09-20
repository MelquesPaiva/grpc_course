package main

import (
	"log"

	pb "github.com/MelquesPaiva/grpc-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	addr = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Fail to connect: %v\n", err)
	}
	defer conn.Close()

	svc := pb.NewGreetServiceClient(conn)
	// doGreet(svc)
	doGreetManyTimes(svc)
}
