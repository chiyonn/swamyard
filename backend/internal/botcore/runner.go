package botcore

import (
	"context"
	"time"

	pb "github.com/chiyonn/swarmyard/api/proto/pricefeed"
	executorpb "github.com/chiyonn/swarmyard/api/proto/tradeexecutor"
	"github.com/chiyonn/swarmyard/internal/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (b *Bot) Run() {
	var conn *grpc.ClientConn
	var err error

	for i := 0; i < 10; i++ {
		conn, err = grpc.NewClient("pricefeed:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err == nil {
			break
		}
		logger.Error("Retrying gRPC connection to pricefeed... (%d/10): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		logger.Error("Failed to connect to gRPC after retries: %v", err)
		return
	}
	defer conn.Close()

	client := pb.NewPriceFeedClient(conn)

	var stream pb.PriceFeed_SubscribePricesClient

	for i := 0; i < 10; i++ {
		stream, err = client.SubscribePrices(context.Background(), &pb.PriceRequest{Base: b.Pair})
		if err == nil {
			break
		}
		logger.Error("Retrying SubscribePrices... (%d/10): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		logger.Error("Failed to subscribe to price feed after retries: %v", err)
		return
	}

	for {
		snapshot, err := stream.Recv()
		if err != nil {
			logger.Error("Stream error: %v", err)
			break
		}

		for _, rate := range snapshot.Rates {
			action := b.Strategy.Decide(rate.Price)
			if action != "HOLD" {
				go placeOrder(b.ID, rate.Pair, action)
			}
		}
	}
}

func placeOrder(botID, pair, action string) {
	conn, err := grpc.NewClient("executor:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("Failed to connect to gRPC: %v", err)
		return
	}
	defer conn.Close()

	client := executorpb.NewTradeExecutorClient(conn)
	side := executorpb.Side_BUY
	if action == "SELL" {
		side = executorpb.Side_SELL
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := client.PlaceOrder(ctx, &executorpb.OrderRequest{
		BotId:  botID,
		Pair:   pair,
		Amount: 1000,
		Side:   side,
	})
	if err != nil {
		logger.Error("Order failed: %v", err)
		return
	}
	logger.Info("Order executed: id=%s status=%s", res.OrderId, res.Status.String())
}
