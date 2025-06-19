package game

import (
	"fmt"

	"tictactoe.ngyngcphu.com/pkg/types"
)

type Board struct {
	rows    int
	columns int
	grid    [][]types.Symbol
}

func NewBoard(rows, columns int) *Board {
	board := &Board{
		rows:    rows,
		columns: columns,
		grid:    make([][]types.Symbol, rows),
	}
	for i := range rows {
		board.grid[i] = make([]types.Symbol, columns)
		for j := range columns {
			board.grid[i][j] = types.Empty
		}
	}
	return board
}

func (board *Board) GetRows() int {
	return board.rows
}

func (board *Board) GetColumns() int {
	return board.columns
}

func (board *Board) IsValidMove(pos *types.Position) bool {
	return pos.Row >= 0 && pos.Row < board.rows && pos.Col >= 0 && pos.Col < board.columns && board.grid[pos.Row][pos.Col] == types.Empty
}

func (board *Board) MakeMove(pos *types.Position, symbol types.Symbol) {
	board.grid[pos.Row][pos.Col] = symbol
}

func (board *Board) CheckWin() (bool, types.Symbol) {
	for i := range board.rows {
		if board.grid[i][0] != types.Empty && board.isWinningLine(board.grid[i]) {
			return true, board.grid[i][0]
		}
	}

	for i := range board.columns {
		column := make([]types.Symbol, board.rows)
		for j := range board.rows {
			column[j] = board.grid[j][i]
		}
		if column[0] != types.Empty && board.isWinningLine(column) {
			return true, column[0]
		}
	}

	if board.rows == board.columns {
		mainDiagonal := make([]types.Symbol, board.rows)
		for i := range board.rows {
			mainDiagonal[i] = board.grid[i][i]
		}
		if mainDiagonal[0] != types.Empty && board.isWinningLine(mainDiagonal) {
			return true, mainDiagonal[0]
		}

		antiDiagonal := make([]types.Symbol, board.rows)
		for i := range board.rows {
			antiDiagonal[i] = board.grid[i][board.rows-i-1]
		}
		if antiDiagonal[0] != types.Empty && board.isWinningLine(antiDiagonal) {
			return true, antiDiagonal[0]
		}
	}

	return false, types.Empty
}

func (board *Board) PrintBoard() {
	for i := range board.rows {
		for j := range board.columns {
			fmt.Printf(" %s ", board.grid[i][j])
			if j < board.columns-1 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < board.rows-1 {
			for k := range board.columns {
				fmt.Print("---")
				if k < board.columns-1 {
					fmt.Print("+")
				}
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (board *Board) IsFull() bool {
	for i := range board.rows {
		for j := range board.columns {
			if board.grid[i][j] == types.Empty {
				return false
			}
		}
	}
	return true
}

func (board *Board) isWinningLine(line []types.Symbol) bool {
	if len(line) == 0 {
		return false
	}
	first := line[0]
	for _, symbol := range line {
		if symbol != first {
			return false
		}
	}
	return true
}
