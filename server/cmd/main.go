package main

import (
	"grpc-pilot/generated"
	"grpc-pilot/server"
	"log"
	"net"

	"google.golang.org/grpc"
)

var serverAddr = "0.0.0.0:5000"

func main() {
	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatal("server failed to listen on address " + serverAddr + ". Error: " + err.Error())
	}

	log.Print("listening on " + serverAddr)

	grpcServer := grpc.NewServer()

	var serverInstance = server.Server{}
	generated.RegisterSumServiceServer(grpcServer, &serverInstance)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("server failed to serve request with error " + err.Error())
	}
}
