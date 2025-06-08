package handler

import (
	"net/http"
	"time"

	"github.com/chiyonn/swarmyard/internal/logger"
	"github.com/chiyonn/swarmyard/internal/pricefeed/model"
	"github.com/chiyonn/swarmyard/internal/pricefeed/service"

	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	store *service.RateStore
}

func NewWebSocketHandler(store *service.RateStore) *WebSocketHandler {
	return &WebSocketHandler{
		store: store,
	}
}

func (h *WebSocketHandler) GetExchangeRate(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error("WebSocket upgrade failed: %v", err)
		http.Error(w, "WebSocket upgrade failed", http.StatusBadRequest)
		return
	}
	defer conn.Close()

	for {
		v := h.store.Get()
		if v == nil {
			time.Sleep(1 * time.Second)
			continue
		}

		conn.WriteJSON(model.ExchangeRateResponse(*v))
		time.Sleep(1 * time.Second)
	}
}
