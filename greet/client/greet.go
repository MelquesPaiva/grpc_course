package main

import (
	"context"
	"log"
	"time"

	pb "github.com/MelquesPaiva/grpc-course/greet/proto"
)

func doGreet(svc pb.GreetServiceClient) {
	log.Println("doGreet was invoke")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := svc.Greet(ctx, &pb.GreetRequest{
		FirstName: "Melques Paiva",
	})
	if err != nil {
		log.Fatalf("Fail to get a response: %v\n", err)
	}

	log.Printf("%s\n", response.Result)
}
