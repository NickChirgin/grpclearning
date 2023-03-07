package main

import (
	"context"
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
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Nick",
			LastName:  "Dont care",
		},
	}
	res, e := c.Greet(context.Background(), req)
	if e != nil {
		fmt.Println(e)
	} 
	fmt.Println(res.Result)
}