package server

import (
	"context"
	"errors"
	"grpc-pilot/generated"
	"io"
	"log"
	"sync"

	"google.golang.org/grpc"
)

func (s *Server) Sum(ctx context.Context, request *generated.SumRequest) (*generated.SumResponse, error) {
	if request == nil {
		log.Fatal("Invalid request")
	}

	log.Print("sum method invoked")

	return &generated.SumResponse{
		Sum: request.First + request.Second,
	}, nil
}

func getPrimeDecomposition(number int32) ([]int32, error) {
	if number < 2 {
		return nil, errors.New("invalid number provided")
	}

	var primes []int32
	var start int32 = 2

	for {
		if number <= 1 {
			break
		}

		if number%start == 0 {
			primes = append(primes, start)
			number /= start
		} else {
			start += 1
		}
	}

	return primes, nil
}

func (s *Server) PrimeNumberDecomposition(request *generated.PrimeRequest, stream grpc.ServerStreamingServer[generated.PrimeResponse]) error {
	log.Print("PrimeNumberDecomposition method invoked")

	if request == nil {
		log.Fatal("invalid request")
	}

	primesList, err := getPrimeDecomposition(request.Number)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	for _, num := range primesList {
		stream.Send(&generated.PrimeResponse{
			Prime: num,
		})
	}

	return nil
}

func (s *Server) Average(reqStream grpc.ClientStreamingServer[generated.AverageRequest, generated.AverageResponse]) error {
	log.Print("average method invoked")

	var summation = 0.0
	var iterations = 0.0
	for {
		req, err := reqStream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		summation += float64(req.Number)
		iterations += 1.0
	}

	return reqStream.SendAndClose(&generated.AverageResponse{
		Average: float32(summation / iterations),
	})
}

func (s *Server) Max(reqStream grpc.BidiStreamingServer[generated.MaxRequest, generated.MaxResponse]) error {
	var numChannel = make(chan int32, 2)

	var wg = sync.WaitGroup{}

	// go routine to catch stream requests from client
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		for {
			request, err := reqStream.Recv()
			if err != nil {
				if err == io.EOF {
					numChannel <- -1 // this is termination of request
					break
				}

				log.Fatal(err.Error())
			}
			log.Print("request received from client: ", request.Number)

			numChannel <- request.Number
		}

		wg.Done()
	}(&wg)

	// go routine to stream responses to client
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		var maximus int32 = 0
		for {
			var num = <-numChannel

			if num == -1 {
				break
			}

			if num > maximus {
				maximus = num

				// send only when there is a new maxima encountered
				err := reqStream.Send(&generated.MaxResponse{
					MaxNumber: maximus,
				})

				if err != nil {
					log.Fatal(err.Error())
				}
			}
		}

		wg.Done()
	}(&wg)

	wg.Wait()

	return nil
}
