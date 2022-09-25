package main

import (
	"io"
	"log"

	pb "github.com/MelquesPaiva/grpc-course/calculator/proto"
)

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Println("Average function was invoked")
	var sum int64
	var count int64

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Response: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading server response: %v", err)
		}

		sum += req.Value
		count++
	}
}
