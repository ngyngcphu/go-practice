package states

import (
	"testing"

	"tictactoe.ngyngcphu.com/pkg/types"
)

func TestNewGameContext(t *testing.T) {
	context := NewGameContext()

	if _, ok := context.GetCurrentState().(*XTurnState); !ok {
		t.Fatal("NewGameContext should start with XTurnState")
	}

	if context.IsGameOver() {
		t.Fatal("NewGameContext should not be game over initially")
	}
}

func TestGameContextSetState(t *testing.T) {
	context := NewGameContext()
	newState := &OTurnState{}

	context.SetState(newState)
	if context.GetCurrentState() != newState {
		t.Fatal("SetState didnlt update current state")
	}
}

func TestGameContextNext(t *testing.T) {
	context := NewGameContext()
	player := &mockPlayer{symbol: types.X}

	context.Next(player, false)
	if _, ok := context.GetCurrentState().(*OTurnState); !ok {
		t.Fatal("Next should transition from XTurnState to OTurnState")
	}
}
