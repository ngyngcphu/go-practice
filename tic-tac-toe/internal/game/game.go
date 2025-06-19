package game

import (
	"fmt"

	"tictactoe.ngyngcphu.com/internal/states"
	"tictactoe.ngyngcphu.com/internal/strategy"
	"tictactoe.ngyngcphu.com/pkg/types"
)

type BoardGame interface {
	Player()
}

type TicTacToeGame struct {
	board         *Board
	playerX       *Player
	playerO       *Player
	currentPlayer *Player
	gameContext   *states.GameContext
}

func NewTicTacToeGame(xStrategy, oStrategy strategy.PlayerStrategy, rows, columns int) *TicTacToeGame {
	board := NewBoard(rows, columns)
	playerX := NewPlayer(types.X, xStrategy)
	playerO := NewPlayer(types.O, oStrategy)
	gameContext := states.NewGameContext()

	return &TicTacToeGame{
		board:         board,
		playerX:       playerX,
		playerO:       playerO,
		currentPlayer: playerX,
		gameContext:   gameContext,
	}
}

func (game *TicTacToeGame) Play() {
	for !game.gameContext.IsGameOver() {
		game.board.PrintBoard()

		move := game.currentPlayer.GetStrategy().MakeMove(game.board)
		game.board.MakeMove(&move, game.currentPlayer.GetSymbol())

		hasWon, _ := game.board.CheckWin()
		if !hasWon && game.board.IsFull() {
			game.gameContext.SetState(&states.DrawState{})
			break
		}

		game.gameContext.Next(game.currentPlayer, hasWon)
		if !game.gameContext.IsGameOver() {
			game.switchPlayer()
		}
	}
	game.announceResult()
}

func (game *TicTacToeGame) switchPlayer() {
	if game.currentPlayer == game.playerX {
		game.currentPlayer = game.playerO
	} else {
		game.currentPlayer = game.playerX
	}
}

func (game *TicTacToeGame) announceResult() {
	game.board.PrintBoard()

	switch game.gameContext.GetCurrentState().(type) {
	case *states.XWonState:
		fmt.Println("Player X wins!")
	case *states.OWonState:
		fmt.Println("Player O wins!")
	case *states.DrawState:
		fmt.Println("It's a draw!")
	default:
		fmt.Println("Game ended in an unexpected state")
	}
}
