package main

import "fmt"

func main() {
	trafficLight := NewTrafficLightContext()
	fmt.Printf("Initial state: %s\n\n", trafficLight.GetColor())

	trafficLight.Next()
	trafficLight.Next()
	trafficLight.Next()
	trafficLight.Next()

	fmt.Printf("\nCurrent state: %s\n", trafficLight.GetColor())
}
