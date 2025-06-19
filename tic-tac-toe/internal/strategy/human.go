package strategy

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"tictactoe.ngyngcphu.com/pkg/types"
)

type HumanPlayerStrategy struct {
	playerName string
	scanner    *bufio.Scanner
}

func NewHumanPlayerStrategy(playerName string) *HumanPlayerStrategy {
	return &HumanPlayerStrategy{
		playerName: playerName,
		scanner:    bufio.NewScanner(os.Stdin),
	}
}

func (human *HumanPlayerStrategy) GetName() string {
	return human.playerName
}

func (human *HumanPlayerStrategy) MakeMove(board BoardInterface) types.Position {
	for {
		fmt.Printf("%s enter your move (row [0-%d] and column [0-%d]): ", human.playerName, board.GetRows()-1, board.GetColumns()-1)
		if !human.scanner.Scan() {
			continue
		}

		input := human.scanner.Text()
		parts := strings.Fields(input)

		if len(parts) != 2 {
			fmt.Println("Invalid input. Please enter row and column as numbers")
			continue
		}

		row, err1 := strconv.Atoi(parts[0])
		col, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil {
			fmt.Println("Invalid input. Please enter row and column as numbers")
			continue
		}

		move := types.Position{Row: row, Col: col}
		if board.IsValidMove(&move) {
			return move
		}
		fmt.Println("Invalid move. Please try again!")
	}
}
