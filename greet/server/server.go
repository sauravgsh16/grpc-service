package main

import (
	"context"
	"log"
	"net"

	"github.com/sauravgsh16/grpc-tut/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Greet(ctx context.Context, req *greet.GreetRequest) (*greet.GreetResponse, error) {
	firstname := req.GetGreeting().GetFirstname()
	result := "Hello " + firstname
	res := &greet.GreetResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	ln, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error creating listener: %v", err)
	}

	s := grpc.NewServer()

	greet.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(ln); err != nil {
		log.Fatalf("failed to server %v", err)
	}

}
