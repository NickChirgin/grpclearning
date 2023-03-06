package main

import (
	"fmt"
	"log"

	"github.com/nickchirgin/grpclearning/greet/greetpb/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello from client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Can't connect %v", err)
	}
	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)
	fmt.Println(c)
}