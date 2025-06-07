package main

import (
    "context"
    "testing"

    pb "github.com/chiyonn/swarmyard/api/proto/tradeexecutor"
)

func TestExecutorServerPlaceOrder(t *testing.T) {
    s := &ExecutorServer{}
    req := &pb.OrderRequest{BotId: "bot", Pair: "USD/JPY", Amount: 1, Side: pb.Side_BUY}
    resp, err := s.PlaceOrder(context.Background(), req)
    if err != nil {
        t.Fatalf("PlaceOrder returned error: %v", err)
    }
    if resp.Status != pb.OrderStatus_SUCCESS {
        t.Errorf("expected SUCCESS got %v", resp.Status)
    }
    if resp.OrderId == "" {
        t.Errorf("expected non-empty order id")
    }
}

func TestExecutorServerPauseResume(t *testing.T) {
    s := &ExecutorServer{}
    botReq := &pb.BotRequest{BotId: "bot"}
    pause, err := s.PauseBot(context.Background(), botReq)
    if err != nil {
        t.Fatalf("PauseBot error: %v", err)
    }
    if pause.Status != "paused" {
        t.Errorf("expected paused got %s", pause.Status)
    }
    resume, err := s.ResumeBot(context.Background(), botReq)
    if err != nil {
        t.Fatalf("ResumeBot error: %v", err)
    }
    if resume.Status != "resumed" {
        t.Errorf("expected resumed got %s", resume.Status)
    }
}

