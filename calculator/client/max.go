package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/MelquesPaiva/grpc-course/calculator/proto"
)

func max(svc pb.CalculatorServiceClient) {
	log.Println("max was invoke")

	stream, err := svc.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while calling max: %v", err)
	}

	reqs := []*pb.MaxRequest{
		{Value: 1},
		{Value: 5},
		{Value: 3},
		{Value: 6},
		{Value: 2},
		{Value: 20},
	}

	// First go routine send all the requests to the server
	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send a request: %v\n", req)
			stream.Send(req)
			time.Sleep(time.Second * 1)
		}
		stream.CloseSend()
	}()

	// Second go routine get all the responses from the server
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Erro while receiving: %v\n", err)
				break
			}

			log.Printf("Received: %v\n", res)
		}

		close(waitc)
	}()

	// Waiting to the close call
	<-waitc
}
