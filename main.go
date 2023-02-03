package main

import (
	"github.com/seungbohong94/go-grpc-example/tutorial"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	println("gRPC server tutorial in Go")

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	tutorial.RegisterTutorialServer(s, &tutorial.Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
