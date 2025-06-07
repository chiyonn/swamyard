package config

import (
    "os"
    "testing"
)

func TestLoad(t *testing.T) {
    os.Setenv("GRPC_ADDRESS", ":12345")
    os.Setenv("MODE", "DEMO")
    os.Setenv("BOT_ID", "test-bot")
    defer os.Unsetenv("GRPC_ADDRESS")
    defer os.Unsetenv("MODE")
    defer os.Unsetenv("BOT_ID")

    cfg, err := Load()
    if err != nil {
        t.Fatalf("Load returned error: %v", err)
    }
    if cfg.GRPCAddress != ":12345" {
        t.Errorf("GRPCAddress expected :12345 got %s", cfg.GRPCAddress)
    }
    if cfg.Mode != "DEMO" {
        t.Errorf("Mode expected DEMO got %s", cfg.Mode)
    }
    if cfg.BotID != "test-bot" {
        t.Errorf("BotID expected test-bot got %s", cfg.BotID)
    }
}
