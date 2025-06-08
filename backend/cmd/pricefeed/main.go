package main

import (
	"context"
	"net/http"


	"github.com/chiyonn/swarmyard/internal/config"
	"github.com/chiyonn/swarmyard/internal/logger"
	"github.com/chiyonn/swarmyard/internal/pricefeed/router"
	"github.com/chiyonn/swarmyard/internal/pricefeed/service"
	grpcserver "github.com/chiyonn/swarmyard/internal/pricefeed/grpc"

)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	logger.InitLogger()
	logger.Info("PriceFeed starting in %s mode", cfg.Mode)

	store := service.NewRateStore()
	client := service.NewRateClient(store)
	go client.StartPolling(context.Background(), "JPY")

	go func() {
		httpRouter := router.NewRouter(store)
		if err := http.ListenAndServe("0.0.0.0:8080", httpRouter); err != nil {
			logger.Fatal("HTTP server failed: %v", err)
		}
	}()

	gs, err := grpcserver.New()
	if err != nil {
		logger.Fatal("failed to create grpc server: %v", err)
	}
	if err := gs.Start(); err != nil {
		logger.Fatal("failed to start grpc server: %v", err)
	}
}
