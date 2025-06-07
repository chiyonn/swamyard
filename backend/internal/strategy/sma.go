package strategy

type SMAStrategy struct{}

func (s *SMAStrategy) Decide(price float64) string {
	if price <= 140.20 {
		return "BUY"
	} else if price >= 140.80 {
		return "SELL"
	}
	return "HOLD"
}

