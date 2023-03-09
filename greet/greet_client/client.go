package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	doServerStreaming(c)
	doClientStreaming(c)
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Client streaming \n")
	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Nick",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Second",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Third",
			},
		},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling LongGreet: %v", err)
	}
	for _, req := range requests {
		stream.Send(req)
		time.Sleep(time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while recieving response")
	}
	fmt.Printf("Long Response %v\n", res)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("streaming \n")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Nick",
			LastName: "C",
		},
	}
	res, err := c.GreetManyTimes(context.Background(), req)	
	if err != nil {
		log.Fatalf("Error while calling stream rpc %v", err)
	}
	for {
		msg, e := res.Recv()	
		if e == io.EOF {
			break;
		}
		if e != nil {
			log.Fatalf("error while reading stream %v", e)
		}
		fmt.Println(msg.Result)
	}	
}