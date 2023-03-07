package main

import (
	"fmt"
	"log"
	"net"

	"github.com/nickchirgin/grpclearning/calculator/calculatorpb"
	"google.golang.org/grpc"
)
type server struct{
	calculatorpb.UnimplementedPrimeServiceServer
}

func (s *server) Prime(req *calculatorpb.PrimeRequest, stream calculatorpb.PrimeService_PrimeServer) error{
	fmt.Printf("Prime function was invoked with %v\n", req)
	number := req.GetNumber()
	k := int32(2)
	for number > 1 {
		if number % k == 0 {
			res := &calculatorpb.PrimeResponse{
				Prime: k,
			}
			stream.Send(res)
			number = number / k
		} else {
			k++
		}
	}
	return nil
}
func main() {
	fmt.Println("Hello")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterPrimeServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}