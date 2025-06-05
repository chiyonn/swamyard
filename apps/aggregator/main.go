package main

import (
	"github.com/chiyonn/swarmyard/pkg/config"
	"github.com/chiyonn/swarmyard/pkg/logger"
)

func main() {
	_, err := config.Load()
	if err != nil {
		panic(err)
	}

	logger.InitLogger()
	logger.Info("Aggregator service starting...")
}
