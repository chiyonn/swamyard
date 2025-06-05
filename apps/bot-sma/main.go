package main

import (
	"context"
	"net/http"
	"time"

	pricefeedpb "github.com/chiyonn/swarmyard/api/proto/pricefeed"
	pb "github.com/chiyonn/swarmyard/api/proto/tradeexecutor"
	"github.com/chiyonn/swarmyard/pkg/config"
	"github.com/chiyonn/swarmyard/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func startPriceStream() {
	conn, err := grpc.NewClient("pricefeed:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("Failed to connect to PriceFeed: %v", err)
		return
	}
	defer conn.Close()

	client := pricefeedpb.NewPriceFeedClient(conn)
	stream, err := client.SubscribePrices(context.Background(), &pricefeedpb.PriceRequest{Pair: "BTC/USDT"})
	if err != nil {
		logger.Error("Failed to subscribe: %v", err)
		return
	}

	for {
		snapshot, err := stream.Recv()
		if err != nil {
			logger.Error("Failed to receive snapshot: %v", err)
			break
		}
		logger.Info("Received price: %.2f for %s at %d", snapshot.Price, snapshot.Pair, snapshot.Timestamp)
	}
}

func placeOrder() {
	conn, err := grpc.NewClient(
		"executor:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Error("Failed to connect to Executor: %v", err)
		return
	}
	defer conn.Close()

	client := pb.NewTradeExecutorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &pb.OrderRequest{
		BotId:  "bot-sma",
		Pair:   "BTC/USDT",
		Amount: 0.01,
		Side:   pb.Side_BUY,
	}

	res, err := client.PlaceOrder(ctx, req)
	if err != nil {
		logger.Error("PlaceOrder failed: %v", err)
		return
	}

	logger.Info("Order placed: id=%s status=%s msg=%s", res.OrderId, res.Status.String(), res.Message)
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	logger.InitLogger()
	logger.Info("Bot SMA starting in %s mode", cfg.Mode)

	http.HandleFunc("/health", healthHandler)
	go http.ListenAndServe(":8080", nil)

	go startPriceStream()
	go placeOrder()

	select {}
}
