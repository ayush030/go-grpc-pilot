package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

var serverAddr = "0.0.0.0:5000"

func main() {
	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatal("server failed to listen on address " + serverAddr + ". Error: " + err.Error())
	}

	server := grpc.NewServer()

	if err := server.Serve(listener); err != nil {
		log.Fatal("server failed to serve request with error " + err.Error())
	}
}
