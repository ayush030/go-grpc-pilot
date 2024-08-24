package main

import (
	"grpc-pilot/client"
	"log"
	"os"

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

	var argsWithoutProgram = os.Args[1:]
	if len(argsWithoutProgram) == 0 {
		log.Println("Lets do something random")
		argsWithoutProgram = append(argsWithoutProgram, "random")
	}

	switch argsWithoutProgram[0] {
	case "sum":
		if len(argsWithoutProgram) < 3 {
			log.Fatal("invalid inputs, provide atleast 2 integers")
		}

		client.Sum(grpcConnection, argsWithoutProgram[1:]...)

	case "primes":
		if len(argsWithoutProgram) != 2 {
			log.Fatal("invalid inputs, provide only 1 integer")
		}

		client.Primes(grpcConnection, argsWithoutProgram[1])

	case "average":
		if len(argsWithoutProgram) < 2 {
			log.Fatal("invalid inputs, provide atleast 1 number")
		}

		client.Average(grpcConnection, argsWithoutProgram[1:]...)

	case "max":
		if len(argsWithoutProgram) < 2 {
			log.Fatal("invalid inputs, provide atleast 1 number")
		}

		client.Max(grpcConnection, argsWithoutProgram[1:]...)

	case "random":

	}
}
