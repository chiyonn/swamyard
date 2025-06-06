package main

import (
	"github.com/chiyonn/swarmyard/internal/config"
	"github.com/chiyonn/swarmyard/internal/logger"
)

func main() {
	_, err := config.Load()
	if err != nil {
		panic(err)
	}

	logger.InitLogger()
	logger.Info("Aggregator service starting...")
}
