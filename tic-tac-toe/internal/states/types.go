package states

import "tictactoe.ngyngcphu.com/pkg/types"

type GameState interface {
	Next(gc *GameContext, player PlayerInterface, hasWon bool)
	IsGameOver() bool
}

type PlayerInterface interface {
	GetSymbol() types.Symbol
}
