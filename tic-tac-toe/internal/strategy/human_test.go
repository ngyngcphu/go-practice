package strategy

import (
	"bufio"
	"strings"
	"testing"

	"tictactoe.ngyngcphu.com/pkg/types"
)

type fakeBoard struct {
	rows, cols int
	validMoves map[types.Position]bool
}

func (b *fakeBoard) GetRows() int {
	return b.rows
}

func (b *fakeBoard) GetColumns() int {
	return b.cols
}

func (b *fakeBoard) IsValidMove(pos *types.Position) bool {
	if b.validMoves == nil {
		return pos.Row >= 0 && pos.Row < b.rows && pos.Col >= 0 && pos.Col < b.cols
	}
	return b.validMoves[*pos]
}

func TestGetName(t *testing.T) {
	h := NewHumanPlayerStrategy("Alice")
	if got := h.GetName(); got != "Alice" {
		t.Fatalf("GetName() = %q; want %q", got, "Alice")
	}
}

func TestMakeMove_WithInvalidThenValidInput(t *testing.T) {
	board := &fakeBoard{
		rows: 3,
		cols: 3,
		validMoves: map[types.Position]bool{
			{Row: 1, Col: 2}: true,
		},
	}

	input := strings.Join([]string{
		"foo",
		"1 bar",
		"9 9",
		"1 2",
	}, "\n") + "\n"

	h := &HumanPlayerStrategy{
		playerName: "Bob",
		scanner:    bufio.NewScanner(strings.NewReader(input)),
	}

	move := h.MakeMove(board)
	want := types.Position{Row: 1, Col: 2}

	if move != want {
		t.Fatalf("MakeMove() = %+v; want %+v", move, want)
	}
}

func TestMakeMove_AllInBounds_NoExplicitValidMap(t *testing.T) {
	input := "2 1\n"
	h := &HumanPlayerStrategy{
		playerName: "Carol",
		scanner:    bufio.NewScanner(strings.NewReader(input)),
	}

	board := &fakeBoard{
		rows: 4, cols: 4,
	}
	move := h.MakeMove(board)
	want := types.Position{Row: 2, Col: 1}

	if move != want {
		t.Fatalf("MakeMove() = %+v; want %+v", move, want)
	}
}
