package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var commAddr = "localhost:5000"

func main() {
	connection, err := grpc.NewClient(commAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("client failed to connect with error " + err.Error())
	}

	defer func() {
		err = connection.Close()
		if err != nil {
			log.Fatal("connection to address " + commAddr + " failed with error " + err.Error())
		}
	}()
}
