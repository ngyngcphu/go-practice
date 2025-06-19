package tictactoe

import (
	"fmt"

	"tictactoe.ngyngcphu.com/internal/game"
	"tictactoe.ngyngcphu.com/internal/strategy"
)

func Run() {
	fmt.Println("Welcome to Tic-Tac-Toe!")
	fmt.Println("======================")

	playerXStrategy := strategy.NewHumanPlayerStrategy("PlayerX")
	playerOStrategy := strategy.NewHumanPlayerStrategy("PlayerO")

	gameInstance := game.NewTicTacToeGame(playerXStrategy, playerOStrategy, 3, 3)
	gameInstance.Play()
}
