package client

import (
	"context"
	"grpc-pilot/generated"
	"io"
	"log"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func Sum(grpcConnection *grpc.ClientConn, numbers ...string) {
	var sumConnection = generated.NewSumServiceClient(grpcConnection)

	first, err := strconv.ParseInt(numbers[0], 10, 32)
	if err != nil {
		log.Fatal(err.Error())
	}

	second, err := strconv.ParseInt(numbers[1], 10, 32)
	if err != nil {
		log.Fatal(err.Error())
	}

	response, err := sumConnection.Sum(context.Background(), &generated.SumRequest{
		First:  int32(first),
		Second: int32(second),
	})

	if err != nil {
		log.Fatal("unable to connect to server. Error: " + err.Error())
	}

	log.Println("response ", response.Sum)
}

func Primes(grpcConnection *grpc.ClientConn, number string) {
	var primesConnection = generated.NewPrimeNumberDecompositionServiceClient(grpcConnection)

	num, err := strconv.ParseInt(number, 10, 32)
	if err != nil {
		log.Fatal(err.Error())
	}

	respStream, err := primesConnection.PrimeNumberDecomposition(context.Background(), &generated.PrimeRequest{
		Number: int32(num),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	var iter = 1
	for {
		recieved, err := respStream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatal(err.Error())
		}

		log.Println(iter, recieved.Prime)
		iter += 1
	}
}

func Average(grpcConnection *grpc.ClientConn, numbers ...string) {
	var averageConnection = generated.NewAverageServiceClient(grpcConnection)

	var numList []float32

	for _, num := range numbers {
		floatingNumber, err := strconv.ParseFloat(num, 32)
		if err != nil {
			log.Fatal(err.Error())
		}

		numList = append(numList, float32(floatingNumber))
	}

	connectionStream, err := averageConnection.Average(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, num := range numList {
		connectionStream.Send(&generated.AverageRequest{
			Number: num,
		})

		time.Sleep(100 * time.Millisecond)
	}

	response, err := connectionStream.CloseAndRecv()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("response ", response.Average)
}

func Max(grpcConnection *grpc.ClientConn, numbers ...string) {
	maxConnection := generated.NewMaxStreamingServiceClient(grpcConnection)

	connStream, err := maxConnection.Max(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}

	var wg = sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		for _, num := range numbers {
			number, err := strconv.ParseInt(num, 10, 32)
			if err != nil {
				log.Fatal(err.Error())
			}

			err = connStream.Send(&generated.MaxRequest{
				Number: int32(number),
			})

			if err != nil {
				log.Fatal(err.Error())
			}
		}

		connStream.CloseSend()
		wg.Done()
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		for {
			response, err := connStream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}

				log.Fatal(err.Error())
			}

			log.Println("response from server: ", response.MaxNumber)
		}

		wg.Done()
	}(&wg)

	wg.Wait()
}
