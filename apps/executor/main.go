package main

import (
	"context"
	"log"
	"net"

	pb "github.com/chiyonn/swarmyard/api/proto/tradeexecutor"
	"google.golang.org/grpc"
)

type ExecutorServer struct {
	pb.UnimplementedTradeExecutorServer
}

func (s *ExecutorServer) PlaceOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
	log.Printf("PlaceOrder called: %+v", req)

	return &pb.OrderResponse{
		OrderId: "order-1234",
		Status:  pb.OrderStatus_SUCCESS,
		Message: "Simulated order placed.",
	}, nil
}

func (s *ExecutorServer) PauseBot(ctx context.Context, req *pb.BotRequest) (*pb.BotResponse, error) {
	log.Printf("PauseBot called: %+v", req)

	return &pb.BotResponse{
		Status:  "paused",
		Message: "Bot has been paused.",
	}, nil
}

func (s *ExecutorServer) ResumeBot(ctx context.Context, req *pb.BotRequest) (*pb.BotResponse, error) {
	log.Printf("ResumeBot called: %+v", req)

	return &pb.BotResponse{
		Status:  "resumed",
		Message: "Bot has been resumed.",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTradeExecutorServer(grpcServer, &ExecutorServer{})

	log.Println("Executor gRPC server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
