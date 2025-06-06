package botcore

import (
	"context"
	"time"

	"github.com/chiyonn/swarmyard/pkg/logger"
	pb "github.com/chiyonn/swarmyard/api/proto/pricefeed"
	executorpb "github.com/chiyonn/swarmyard/api/proto/tradeexecutor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (b *Bot) Run() {
	conn, _ := grpc.NewClient("pricefeed:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	client := pb.NewPriceFeedClient(conn)
	stream, _ := client.SubscribePrices(context.Background(), &pb.PriceRequest{Pair: b.Pair})

	for {
		snapshot, err := stream.Recv()
		if err != nil {
			logger.Error("Stream error: %v", err)
			break
		}

		logger.Info("Received price: %.2f", snapshot.Price)

		action := b.Strategy.Decide(snapshot.Price)
		if action != "HOLD" {
			go placeOrder(b.ID, b.Pair, action)
		}
	}
}

func placeOrder(botID, pair, action string) {
	conn, _ := grpc.NewClient("executor:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
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

