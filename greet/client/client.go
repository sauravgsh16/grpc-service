package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sauravgsh16/grpc-tut/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	defer cc.Close()

	c := greet.NewGreetServiceClient(cc)

	doUnary(c)
}

func doUnary(c greet.GreetServiceClient) {
	fmt.Println("Starting unary connection")

	req := &greet.GreetRequest{
		Greeting: &greet.Greeting{
			Firstname: "Saurav",
			Lastname:  "Ghosh",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling RPC %v", err)
	}
	log.Printf("Response from greet: %v\n", res.Result)
}
