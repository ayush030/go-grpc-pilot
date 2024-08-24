package main

import (
	"context"
	"fmt"
	"grpc-pilot/generated"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var commAddr = "localhost:5000"

func main() {
	grpcConnection, err := grpc.NewClient(commAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("client failed to connect with error " + err.Error())
	}

	defer func() {
		err = grpcConnection.Close()
		if err != nil {
			log.Fatal("connection to address " + commAddr + " failed with error " + err.Error())
		}
	}()

	var sumConnection = generated.NewSumServiceClient(grpcConnection)

	response, err := sumConnection.Sum(context.Background(), &generated.SumRequest{
		First:  2,
		Second: 3,
	})

	if err != nil {
		log.Fatal("unable to connect to server. Error: " + err.Error())
	}

	fmt.Println("response ", response.Sum)
}
