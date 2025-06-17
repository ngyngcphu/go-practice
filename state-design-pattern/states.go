package main

import "fmt"

type TrafficLightState interface {
	Next(context *TrafficLightContext)
	GetColor() string
}

type RedState struct{}

func (r *RedState) Next(context *TrafficLightContext) {
	fmt.Println("🔴 Switching from RED to GREEN. Cars go!")
	context.SetState(&GreenState{})
}

func (r *RedState) GetColor() string {
	return "RED"
}

type GreenState struct{}

func (g *GreenState) Next(context *TrafficLightContext) {
	fmt.Println("🟢 Switching from GREEN to YELLOW. Slow down!")
	context.SetState(&YellowState{})
}

func (g *GreenState) GetColor() string {
	return "GREEN"
}

type YellowState struct{}

func (y *YellowState) Next(context *TrafficLightContext) {
	fmt.Println("🟡 Switching from YELLOW to RED. Stop!")
	context.SetState(&RedState{})
}

func (y *YellowState) GetColor() string {
	return "YELLOW"
}
