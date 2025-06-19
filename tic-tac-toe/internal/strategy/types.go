package strategy

import "tictactoe.ngyngcphu.com/pkg/types"

type PlayerStrategy interface {
	MakeMove(board BoardInterface) types.Position
	GetName() string
}

type BoardInterface interface {
	IsValidMove(pos *types.Position) bool
	GetRows() int
	GetColumns() int
}
