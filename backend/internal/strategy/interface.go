package strategy

type Strategy interface {
	Decide(price float64) (action string) // "BUY", "SELL", "HOLD"
}
