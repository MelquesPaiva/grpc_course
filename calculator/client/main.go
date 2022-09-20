package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/MelquesPaiva/grpc-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	addr = "localhost:50052"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Fail to connect: %v\n", err)
	}
	defer conn.Close()

	svc := pb.NewCalculatorServiceClient(conn)
	primeNumber(svc)
}

func sumValue(svc pb.CalculatorServiceClient) {
	log.Println("sumValue was invoke")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := svc.Sum(ctx, &pb.SumRequest{
		FirstValue:  10,
		SecondValue: 5,
	})
	if err != nil {
		log.Fatalf("A error occured while adding the values: %v", err)
	}

	log.Println("Result of the adition:", res.Result)
}

func primeNumber(svc pb.CalculatorServiceClient) {
	log.Println("primeNumber was invoke")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := svc.PrimeNumber(ctx, &pb.PrimeRequest{
		Value: 120,
	})
	if err != nil {
		log.Fatalf("A error occured while adding the values: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Value: %d\n", msg.Result)
	}
}
