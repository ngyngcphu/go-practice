package game

import (
	"testing"

	"tictactoe.ngyngcphu.com/internal/strategy"
	"tictactoe.ngyngcphu.com/pkg/types"
)

type mockStrategy struct {
	name string
	move types.Position
}

func (m *mockStrategy) MakeMove(board strategy.BoardInterface) types.Position {
	return m.move
}

func (m *mockStrategy) GetName() string {
	return m.name
}

func TestPlayer(t *testing.T) {
	strategy := &mockStrategy{name: "TestPlayer", move: types.Position{Row: 1, Col: 1}}
	player := NewPlayer(types.X, strategy)

	if player.GetSymbol() != types.X {
		t.Fatalf("NewPlayer() symbol = %v, want X", player.GetSymbol())
	}

	if player.GetStrategy() != strategy {
		t.Fatal("NewPlayer() didn't set strategy correctly")
	}
}
