package server

import (
	"grpc-pilot/generated"
)

type Server struct {
	generated.SumServiceServer
	generated.PrimeNumberDecompositionServiceServer
	generated.AverageServiceServer
	generated.MaxStreamingServiceServer
}
