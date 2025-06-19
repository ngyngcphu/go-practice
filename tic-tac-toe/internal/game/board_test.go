package game

import (
	"testing"

	"tictactoe.ngyngcphu.com/pkg/types"
)

func TestNewBoard(t *testing.T) {
	board := NewBoard(3, 3)
	if board.rows != 3 || board.columns != 3 {
		t.Fatalf("NewBoard(3, 3) created board with dimensions %dx%d, want 3x3", board.rows, board.columns)
	}

	for i := range board.rows {
		for j := range board.columns {
			if board.grid[i][j] != types.Empty {
				t.Fatalf("NewBoard() cell (%d, %d) = %v, want Empty", i, j, board.grid[i][j])
			}
		}
	}
}

func TestBoardIsValidMove(t *testing.T) {
	board := NewBoard(3, 3)
	tests := []struct {
		name     string
		pos      types.Position
		expected bool
	}{
		{"Valid move center", types.Position{Row: 1, Col: 1}, true},
		{"Valid move corner", types.Position{Row: 0, Col: 0}, true},
		{"Invalid negative row", types.Position{Row: -1, Col: 0}, false},
		{"Invalid negative col", types.Position{Row: 0, Col: -1}, false},
		{"Invalid row too high", types.Position{Row: 3, Col: 0}, false},
		{"Invalid col too high", types.Position{Row: 0, Col: 3}, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := board.IsValidMove(&test.pos)
			if result != test.expected {
				t.Fatalf("IsValidMove(%v) = %v, want %v", test.pos, result, test.expected)
			}
		})
	}
}

func TestBoardMakeMove(t *testing.T) {
	board := NewBoard(3, 3)
	pos := types.Position{Row: 1, Col: 1}

	board.MakeMove(&pos, types.X)
	if board.grid[1][1] != types.X {
		t.Fatalf("MakeMove didn't place symbol correctly: got %v, want X", board.grid[1][1])
	}

	if board.IsValidMove(&pos) {
		t.Fatalf("IsValidMove should return false for occupied position")
	}
}

func TestBoardCheckWin(t *testing.T) {
	tests := []struct {
		name   string
		setup  func(*Board)
		hasWon bool
		winner types.Symbol
	}{
		{
			name: "Row win",
			setup: func(b *Board) {
				b.MakeMove(&types.Position{Row: 0, Col: 0}, types.X)
				b.MakeMove(&types.Position{Row: 0, Col: 1}, types.X)
				b.MakeMove(&types.Position{Row: 0, Col: 2}, types.X)
			},
			hasWon: true,
			winner: types.X,
		},
		{
			name: "Column win",
			setup: func(b *Board) {
				b.MakeMove(&types.Position{Row: 0, Col: 0}, types.O)
				b.MakeMove(&types.Position{Row: 1, Col: 0}, types.O)
				b.MakeMove(&types.Position{Row: 2, Col: 0}, types.O)
			},
			hasWon: true,
			winner: types.O,
		},
		{
			name: "Diagonal win",
			setup: func(b *Board) {
				b.MakeMove(&types.Position{Row: 0, Col: 0}, types.X)
				b.MakeMove(&types.Position{Row: 1, Col: 1}, types.X)
				b.MakeMove(&types.Position{Row: 2, Col: 2}, types.X)
			},
			hasWon: true,
			winner: types.X,
		},
		{
			name: "Anti-diagonal win",
			setup: func(b *Board) {
				b.MakeMove(&types.Position{Row: 0, Col: 2}, types.O)
				b.MakeMove(&types.Position{Row: 1, Col: 1}, types.O)
				b.MakeMove(&types.Position{Row: 2, Col: 0}, types.O)
			},
			hasWon: true,
			winner: types.O,
		},
		{
			name: "No win",
			setup: func(b *Board) {
				b.MakeMove(&types.Position{Row: 0, Col: 0}, types.X)
				b.MakeMove(&types.Position{Row: 1, Col: 1}, types.O)
			},
			hasWon: false,
			winner: types.Empty,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			board := NewBoard(3, 3)
			test.setup(board)

			hasWon, winner := board.CheckWin()
			if hasWon != test.hasWon {
				t.Fatalf("CheckWin() hasWon = %v, want %v", hasWon, test.hasWon)
			}
			if winner != test.winner {
				t.Fatalf("CheckWin() winner = %v, want %v", winner, test.winner)
			}
		})
	}
}

func TestBoardIsFull(t *testing.T) {
	board := NewBoard(2, 2)
	if board.IsFull() {
		t.Fatal("Empty board should not be full")
	}

	positions := []types.Position{
		{Row: 0, Col: 0},
		{Row: 0, Col: 1},
		{Row: 1, Col: 0},
		{Row: 1, Col: 1},
	}

	for i, pos := range positions {
		symbol := types.X
		if i%2 == 1 {
			symbol = types.O
		}
		board.MakeMove(&pos, symbol)
	}
	if !board.IsFull() {
		t.Fatal("Full board should return true for IsFull()")
	}
}
