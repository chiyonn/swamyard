package config

import "github.com/caarlos0/env/v6"

type Config struct {
	GRPCAddress string `env:"GRPC_ADDRESS" envDefault:":50051"`
	Mode        string `env:"MODE" envDefault:"REAL"`
	BotID       string `env:"BOT_ID"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	return cfg, err
}
