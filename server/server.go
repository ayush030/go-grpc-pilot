package server

import (
	"context"
	"grpc-pilot/generated"
	"log"
)

type Server struct {
	generated.SumServiceServer
}

func (s *Server) Sum(ctx context.Context, request *generated.SumRequest) (*generated.SumResponse, error) {
	if request == nil {
		log.Fatal("Invalid request")
	}

	log.Print("sum method invoked")

	return &generated.SumResponse{
		Sum: request.First + request.Second,
	}, nil
}
