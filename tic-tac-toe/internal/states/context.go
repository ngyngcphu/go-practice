package states

type GameContext struct {
	currentState GameState
}

func NewGameContext() *GameContext {
	return &GameContext{
		currentState: &XTurnState{},
	}
}

func (gc *GameContext) Next(player PlayerInterface, hasWon bool) {
	gc.currentState.Next(gc, player, hasWon)
}

func (gc *GameContext) SetState(state GameState) {
	gc.currentState = state
}

func (gc *GameContext) GetCurrentState() GameState {
	return gc.currentState
}

func (gc *GameContext) IsGameOver() bool {
	return gc.currentState.IsGameOver()
}
