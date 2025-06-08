package service

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/chiyonn/swarmyard/internal/logger"
	"github.com/chiyonn/swarmyard/internal/pricefeed/model"
)

type RateClient struct {
	host  string
	store *RateStore
}

func NewRateClient(store *RateStore) *RateClient {
	return &RateClient{
		host:  "http://192.168.40.107:8080",
		store: store,
	}
}

func (c *RateClient) fetchRate(baseCurrency string) *model.ExchangeRateResponse {
	req, err := http.NewRequest("GET", c.host+"/price", nil)
	if err != nil {
		logger.Info("failed to create request: %v", err)
		return nil
	}
	q := req.URL.Query()
	q.Add("currency", baseCurrency)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Info("failed to fetch rates: %v", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Info("non-OK response status: %d", resp.StatusCode)
		return nil
	}

	var data model.ExchangeRateResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		logger.Info("failed to decode response: %v", err)
		return nil
	}

	return &data
}

func (c *RateClient) StartPolling(ctx context.Context, baseCurrency string) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			snapshot := c.fetchRate(baseCurrency)
			if snapshot != nil {
				c.store.Set(snapshot)
			}
		case <-ctx.Done():
			logger.Info("RateClient polling stopped")
			return
		}
	}
}
