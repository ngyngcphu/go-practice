package states

import (
	"testing"

	"tictactoe.ngyngcphu.com/pkg/types"
)

type mockPlayer struct {
	symbol types.Symbol
}

func (m *mockPlayer) GetSymbol() types.Symbol {
	return m.symbol
}

func TestXTurnState(t *testing.T) {
	state := &XTurnState{}
	context := NewGameContext()
	player := &mockPlayer{symbol: types.X}

	if state.IsGameOver() {
		t.Fatal("XTurnState should not be game over")
	}

	state.Next(context, player, true)
	if _, ok := context.GetCurrentState().(*XWonState); !ok {
		t.Fatal("XTurnState should transition to XWonState when X wins")
	}

	state.Next(context, player, false)
	if _, ok := context.GetCurrentState().(*OTurnState); !ok {
		t.Fatal("XTurnState should transition to OTurnState when no win")
	}
}

func TestOTurnState(t *testing.T) {
	state := &OTurnState{}
	context := NewGameContext()
	player := &mockPlayer{symbol: types.O}

	if state.IsGameOver() {
		t.Fatal("oTurnState should not br game over")
	}

	state.Next(context, player, true)
	if _, ok := context.GetCurrentState().(*OWonState); !ok {
		t.Fatal("OTurnState should transition to OWonState when O wins")
	}
}

func TestXWonState(t *testing.T) {
	state := &XWonState{}

	if !state.IsGameOver() {
		t.Fatal("XWonState should be game over")
	}
}

func TestOWonState(t *testing.T) {
	state := &OWonState{}

	if !state.IsGameOver() {
		t.Fatal("OWonState should be game over")
	}
}

func TestDrawState(t *testing.T) {
	state := &DrawState{}

	if !state.IsGameOver() {
		t.Fatal("DrawState should be game over")
	}
}
