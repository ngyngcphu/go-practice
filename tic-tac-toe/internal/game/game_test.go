package game

import (
	"testing"

	"tictactoe.ngyngcphu.com/internal/states"
	"tictactoe.ngyngcphu.com/internal/strategy"
	"tictactoe.ngyngcphu.com/pkg/types"
)

type testStrategy struct {
	name  string
	moves []types.Position
	index int
}

func (t *testStrategy) GetName() string {
	return t.name
}

func (t *testStrategy) MakeMove(board strategy.BoardInterface) types.Position {
	if t.index >= len(t.moves) {
		t.index = 0
		return types.Position{Row: 0, Col: 0}
	}
	move := t.moves[t.index]
	t.index++
	return move
}

func TestNewTicTacToeGame(t *testing.T) {
	xStrategy := &testStrategy{name: "PlayerX", moves: []types.Position{{Row: 0, Col: 0}}}
	oStrategy := &testStrategy{name: "PlayerO", moves: []types.Position{{Row: 0, Col: 1}}}

	game := NewTicTacToeGame(xStrategy, oStrategy, 3, 3)

	if game.board == nil {
		t.Fatal("NewTicTacToeGame should create board")
	}
	if game.playerX == nil || game.playerO == nil {
		t.Fatal("NewTicTacToeGame should create players")
	}
	if game.currentPlayer != game.playerX {
		t.Fatal("NewTicTacToeGame should start with player X")
	}
}

func TestTicTacToeGamePlayerXWins(t *testing.T) {
	xMoves := []types.Position{
		{Row: 0, Col: 0},
		{Row: 1, Col: 1},
		{Row: 2, Col: 2},
	}
	oMoves := []types.Position{
		{Row: 0, Col: 1},
		{Row: 0, Col: 2},
	}

	xStrategy := &testStrategy{name: "PlayerX", moves: xMoves}
	oStrategy := &testStrategy{name: "PlayerO", moves: oMoves}

	game := NewTicTacToeGame(xStrategy, oStrategy, 3, 3)
	game.Play()

	state := game.gameContext.GetCurrentState()
	if _, ok := state.(*states.XWonState); !ok {
		t.Fatalf("expected game state to be *states.XWonState, got %T", state)
	}
}
