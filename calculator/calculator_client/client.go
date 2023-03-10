package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	bidiStreaming(c)
}
func bidiStreaming(c calculatorpb.PrimeServiceClient) {
	fmt.Println("Client streaming")
	waitCh := make(chan struct{})
	stream, err := c.MaxInteger(context.Background())	
	if err != nil {
		log.Fatalf("Error while reading stream %v", err)
	}
	requests := []*calculatorpb.MaxIntegerRequest{
		&calculatorpb.MaxIntegerRequest{Number: 1},
		&calculatorpb.MaxIntegerRequest{Number: 5},
		&calculatorpb.MaxIntegerRequest{Number: 3},
		&calculatorpb.MaxIntegerRequest{Number: 18},
		&calculatorpb.MaxIntegerRequest{Number: 3},
	}
	go func() {
		for _, req := range requests {
			fmt.Printf("Stream sending %v", req)
			stream.Send(req)
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()
	go func() {
		for {
			res, err := stream.Recv()	
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while reading stream %v", err)
				break
			}
			fmt.Printf("Max integer is %d\n", res.GetMax())
		}
		close(waitCh)
	}()
	<-waitCh
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