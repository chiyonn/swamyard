package main

import (
	"github.com/chiyonn/swarmyard/pkg/config"
	"github.com/chiyonn/swarmyard/pkg/logger"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	logger.InitLogger()
	logger.Info("PriceFeed starting in %s mode", cfg.Mode)

	http.HandleFunc("/health", healthHandler)
	go http.ListenAndServe(":8080", nil)

	select {}
}
