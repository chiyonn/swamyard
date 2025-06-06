package botcore

import "github.com/chiyonn/swarmyard/pkg/strategy"

type Bot struct {
	ID       string
	Pair     string
	Strategy strategy.Strategy
}

