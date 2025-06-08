package model

type ExchangeRateResponse struct {
	Base  string     `json:"base"`
	Rates []Currency `json:"rates"`
}
