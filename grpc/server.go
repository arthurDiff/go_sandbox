package main

import (
	"fmt"
	"grpc/comm"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	PORT = 9000
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))

	if err != nil {
		log.Fatalf("Failed to listen on port %d with err: %v", PORT, err)
	}

	s := comm.Server{}

	grpcServer := grpc.NewServer()

	comm.RegisterCommServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %d with err: %v", PORT, err)
	}
}
