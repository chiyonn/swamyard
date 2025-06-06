package main

import (
	"context"
	"net/http"
	"time"

	pricefeedpb "github.com/chiyonn/swarmyard/api/proto/pricefeed"
	pb "github.com/chiyonn/swarmyard/api/proto/tradeexecutor"
	"github.com/chiyonn/swarmyard/pkg/botcore"
	"github.com/chiyonn/swarmyard/pkg/config"
	"github.com/chiyonn/swarmyard/pkg/logger"
	"github.com/chiyonn/swarmyard/pkg/strategy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var lastAction string

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
			logger.Error("Stream error: %v", err)
			break
		}

		if snapshot.Price <= 140.20 && lastAction != "BUY" {
			placeOrder(pb.Side_BUY, "USD/JPY", 1000.0)
			lastAction = "BUY"
		} else if snapshot.Price >= 140.80 && lastAction != "SELL" {
			placeOrder(pb.Side_SELL, "USD/JPY", 1000.0)
			lastAction = "SELL"
		}
	}
}

func placeOrder(side pb.Side, pair string, amount float64) {
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
		Pair:   pair,
		Amount: amount,
		Side:   side,
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

	sma := &strategy.SMAStrategy{}
	bot := &botcore.Bot{
		ID: "bot-sma",
		Pair: "USD/JPY",
		Strategy: sma,
	}
	go bot.Run()

	select {}
}
