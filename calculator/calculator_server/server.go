package main

import (
	"fmt"
	"io"
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
func (s *server) MaxInteger(stream calculatorpb.PrimeService_MaxIntegerServer) error {
	fmt.Println("Max integer server on the line")
	m := int64(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while recieving stream %v", err)
			return err
		}
		m = max(m, req.GetNumber())
		err = stream.Send(&calculatorpb.MaxIntegerResponse{
			Max: m,
		})
		if err != nil {
			log.Fatalf("Error while sending max %v", err)
			return err
		}
	}
}
func (*server) ComputeAverage(stream calculatorpb.PrimeService_ComputeAverageServer) error {
	fmt.Println("Compute average server")
	result, count := 0, 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				Average: int32(result/count),
			})
		}
		if err != nil {
			log.Fatalf("Error while reading stream %v", err)
		}
		result += int(req.GetNumber())
		count++
	}
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
func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}