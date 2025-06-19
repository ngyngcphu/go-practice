package game

import (
	"tictactoe.ngyngcphu.com/internal/strategy"
	"tictactoe.ngyngcphu.com/pkg/types"
)

type Player struct {
	symbol   types.Symbol
	strategy strategy.PlayerStrategy
}

func NewPlayer(symbol types.Symbol, strategy strategy.PlayerStrategy) *Player {
	return &Player{
		symbol:   symbol,
		strategy: strategy,
	}
}

func (p *Player) GetSymbol() types.Symbol {
	return p.symbol
}

func (p *Player) GetStrategy() strategy.PlayerStrategy {
	return p.strategy
}
