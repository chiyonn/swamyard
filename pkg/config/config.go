package config

type Config struct {
	Mode string // demo or real
}

func New() *Config {
	return &Config{Mode: "demo"}
}
