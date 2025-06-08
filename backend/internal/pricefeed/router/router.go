package router

import (
	"net/http"

	"github.com/chiyonn/swarmyard/internal/pricefeed/handler"
	"github.com/chiyonn/swarmyard/internal/pricefeed/service"
)

func NewRouter(store *service.RateStore) *http.ServeMux {
	mux := http.NewServeMux()

	wsh := handler.NewWebSocketHandler(store)
	mux.HandleFunc("/health", handler.HealthHandler)
	mux.HandleFunc("/ws/price", wsh.GetExchangeRate)

	return mux
}
