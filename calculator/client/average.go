package main

import (
	"context"
	"log"

	pb "github.com/MelquesPaiva/grpc-course/calculator/proto"
)

func average(svc pb.CalculatorServiceClient) {
	log.Println("average was invoke")

	reqs := []*pb.AverageRequest{
		{Value: 1},
		{Value: 4},
		{Value: 7},
		{Value: 2},
	}

	stream, err := svc.Average(context.Background())
	if err != nil {
		log.Fatalf("Error while calling average: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from average: %v", err)
	}

	log.Printf("Average response: %f", res.Response)

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

}
