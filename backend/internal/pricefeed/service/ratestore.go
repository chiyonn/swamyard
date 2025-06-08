package service

import (
	"sync/atomic"

	"github.com/chiyonn/swarmyard/internal/pricefeed/model"
)

type RateStore struct {
	value atomic.Value // *model.ExchangeRateResponse
}

func NewRateStore() *RateStore {
	store := &RateStore{}
	store.value.Store((*model.ExchangeRateResponse)(nil))
	return store
}

func (s *RateStore) Set(snapshot *model.ExchangeRateResponse) {
	s.value.Store(snapshot)
}

func (s *RateStore) Get() *model.ExchangeRateResponse {
	v := s.value.Load()
	if v == nil {
		return nil
	}
	return v.(*model.ExchangeRateResponse)
}
