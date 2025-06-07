package strategy

import "testing"

func TestSMAStrategyDecide(t *testing.T) {
    cases := []struct {
        price    float64
        expected string
    }{
        {140.10, "BUY"},
        {140.90, "SELL"},
        {140.50, "HOLD"},
    }

    s := &SMAStrategy{}
    for _, c := range cases {
        if got := s.Decide(c.price); got != c.expected {
            t.Errorf("price %.2f expected %s got %s", c.price, c.expected, got)
        }
    }
}
