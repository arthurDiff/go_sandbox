package comm

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SaySomething(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)

	return &Message{Body: "Hello From Serv"}, nil
}

func (s *Server) mustEmbedUnimplementedCommServiceServer() {
	log.Fatalf("Unimplemented Comm Service")
}
