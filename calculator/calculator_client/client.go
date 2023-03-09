package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/nickchirgin/grpclearning/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello from client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Can't connect %v", err)
	}
	defer conn.Close()
	c := calculatorpb.NewPrimeServiceClient(conn)
	doClientStreaming(c)
}
func doClientStreaming(c calculatorpb.PrimeServiceClient) {
	fmt.Println("Client streaming")
	requests := []*calculatorpb.ComputeAverageRequest{
		&calculatorpb.ComputeAverageRequest{Number: 1},
		&calculatorpb.ComputeAverageRequest{Number: 5},
		&calculatorpb.ComputeAverageRequest{Number: 3},
		&calculatorpb.ComputeAverageRequest{Number: 18},
		&calculatorpb.ComputeAverageRequest{Number: 3},
	}
	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while streaming %v", err)
	}
	for _, req := range requests {
		fmt.Printf("Stream sending %v", req)
		stream.Send(req)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Println("Average is %d", res)
}

func doServerStreaming(c calculatorpb.PrimeServiceClient) {
	fmt.Println("streaming \n")
	req := &calculatorpb.PrimeRequest{
		Number: 120,
	}
	res, err := c.Prime(context.Background(), req)	
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
		fmt.Println(msg.Prime)
	}	
}