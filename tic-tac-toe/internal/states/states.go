package states

import "tictactoe.ngyngcphu.com/pkg/types"

type XTurnState struct{}

func (xturn *XTurnState) Next(gc *GameContext, player PlayerInterface, hasWon bool) {
	if hasWon {
		if player.GetSymbol() == types.X {
			gc.SetState(&XWonState{})
		} else {
			gc.SetState(&OWonState{})
		}
	} else {
		gc.SetState(&OTurnState{})
	}
}

func (xturn *XTurnState) IsGameOver() bool {
	return false
}

type OTurnState struct{}

func (oturn *OTurnState) Next(gc *GameContext, player PlayerInterface, hasWon bool) {
	if hasWon {
		if player.GetSymbol() == types.O {
			gc.SetState(&OWonState{})
		} else {
			gc.SetState(&XWonState{})
		}
	} else {
		gc.SetState(&XTurnState{})
	}
}

func (oturn *OTurnState) IsGameOver() bool {
	return false
}

type XWonState struct{}

func (xwon *XWonState) Next(gc *GameContext, player PlayerInterface, hasWon bool) {
}

func (xwon *XWonState) IsGameOver() bool {
	return true
}

type OWonState struct{}

func (owon *OWonState) Next(gc *GameContext, player PlayerInterface, hasWon bool) {
}

func (owon *OWonState) IsGameOver() bool {
	return true
}

type DrawState struct{}

func (draw *DrawState) Next(gc *GameContext, player PlayerInterface, hasWon bool) {
}

func (draw *DrawState) IsGameOver() bool {
	return true
}
