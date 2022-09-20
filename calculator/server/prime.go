package main

import (
	"log"

	pb "github.com/MelquesPaiva/grpc-course/calculator/proto"
)

func (s *Server) PrimeNumber(req *pb.PrimeRequest, stream pb.CalculatorService_PrimeNumberServer) error {
	log.Printf("PrimeNumber was invoked: %v\n", req)

	start := 2
	value := int(req.Value)

	for value > 1 {
		if value%start == 0 {
			value = value / start
			stream.Send(&pb.PrimeResponse{
				Result: int64(start),
			})
			continue
		}

		start = start + 1
	}

	return nil
}
