package main

import (
	"context"
	"io"
	"log"

	pb "github.com/MelquesPaiva/grpc-course/greet/proto"
)

func doGreetManyTimes(svc pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "Melques",
	}

	stream, err := svc.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("GreetManyTime: %s\n", msg.Result)
	}
}
