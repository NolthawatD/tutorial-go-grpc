package main

import (
	"fmt"
	"log"
	"net"
	"server/services"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	services.RegisterCalculatorServer(s, services.NewCaculatorServer())

	fmt.Println("gRPC server listening on prot 50051")
	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
