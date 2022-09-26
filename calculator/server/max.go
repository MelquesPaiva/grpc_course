package main

import (
	"io"
	"log"

	pb "github.com/MelquesPaiva/grpc-course/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max function was invoked")

	var currentMax int64 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		if req.Value > currentMax {
			stream.Send(&pb.MaxResponse{
				Max: req.Value,
			})
			currentMax = req.Value
		}
	}

	return nil
}
