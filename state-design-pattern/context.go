package main

type TrafficLightContext struct {
	currentState TrafficLightState
}

func NewTrafficLightContext() *TrafficLightContext {
	return &TrafficLightContext{
		currentState: &RedState{},
	}
}

func (t *TrafficLightContext) SetState(state TrafficLightState) {
	t.currentState = state
}

func (t *TrafficLightContext) Next() {
	t.currentState.Next(t)
}

func (t *TrafficLightContext) GetColor() string {
	return t.currentState.GetColor()
}
