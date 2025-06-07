package botcore

import "github.com/chiyonn/swarmyard/internal/strategy"

type Bot struct {
	ID       string
	Pair     string
	Strategy strategy.Strategy
}

