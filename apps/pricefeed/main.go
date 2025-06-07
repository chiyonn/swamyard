package main

import (
	"math/rand"
	"net"
	"net/http"
	"sync/atomic"
	"time"

	pb "github.com/chiyonn/swarmyard/api/proto/pricefeed"
	"github.com/chiyonn/swarmyard/pkg/config"
	"github.com/chiyonn/swarmyard/pkg/logger"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

var latestPrice atomic.Value

type PriceFeedServer struct {
	pb.UnimplementedPriceFeedServer
}

func (s *PriceFeedServer) SubscribePrices(req *pb.PriceRequest, stream pb.PriceFeed_SubscribePricesServer) error {
	pair := req.Pair
	logger.Info("New subscription for pair: %s", pair)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			price := 140.00 + rand.Float64() // → 140.00〜140.99
			snapshot := &pb.PriceSnapshot{
				Pair:      pair,
				Price:     price,
				Timestamp: time.Now().Unix(),
			}

			latestPrice.Store(snapshot)

			if err := stream.Send(snapshot); err != nil {
				logger.Error("Failed to send snapshot: %v", err)
				return err
			}

			logger.Info("Sent price: %.2f for %s", price, pair)

		case <-stream.Context().Done():
			logger.Info("Stream closed by client for %s", pair)
			return nil
		}
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func streamHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error("WebSocket upgrade failed: %v", err)
		return
	}
	defer conn.Close()

	for {
		v := latestPrice.Load()
		if v == nil {
			time.Sleep(1 * time.Second)
			continue
		}

		snapshot := v.(*pb.PriceSnapshot)
		conn.WriteJSON(map[string]interface{}{
			"pair":  snapshot.Pair,
			"price": snapshot.Price,
			"time":  snapshot.Timestamp,
		})
		time.Sleep(1 * time.Second)
	}
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	logger.InitLogger()
	logger.Info("PriceFeed starting in %s mode", cfg.Mode)

	// gRPC server setup
	go func() {
		lis, err := net.Listen("tcp", ":50052")
		if err != nil {
			logger.Fatal("Failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		pb.RegisterPriceFeedServer(grpcServer, &PriceFeedServer{})

		logger.Info("gRPC server listening on :50052")
		if err := grpcServer.Serve(lis); err != nil {
			logger.Fatal("Failed to serve: %v", err)
		}
	}()

	// HTTP health check
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/ws/price", streamHandler)
	go http.ListenAndServe(":8080", nil)

	select {}
}
