package main

import (
	"context"
	"fmt"
	"grpc/comm"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const PORT = 9000

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(fmt.Sprintf(":%d", 9000), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to port:%d with err: %v", PORT, err)
	}
	defer conn.Close()

	c := comm.NewCommServiceClient(conn)

	msg := comm.Message{
		Body: "Hello from client",
	}

	response, err := c.SaySomething(context.Background(), &msg)

	if err != nil {
		log.Fatalf("Failed sending msg via grpc SayingSomething err: %v", err)
	}

	log.Printf("Response from Server: %s", response.Body)
}
