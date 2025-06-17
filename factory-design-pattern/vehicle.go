package main

import (
	"fmt"
	"strings"
)

type Vehicle interface {
	Start()
	Stop()
}

type Car struct{}

func (c Car) Start() {
	fmt.Println("Car is starting...")
}

func (c Car) Stop() {
	fmt.Println("Car is stopping...")
}

type Truck struct{}

func (t Truck) Start() {
	fmt.Println("Truck is starting...")
}

func (t Truck) Stop() {
	fmt.Println("Truck is stopping...")
}

type Bike struct{}

func (b Bike) Start() {
	fmt.Println("Bike is starting...")
}

func (b Bike) Stop() {
	fmt.Println("Bike is stopping...")
}

func CreateVehicle(vehicleType string) Vehicle {
	switch strings.ToLower(vehicleType) {
	case "car":
		return Car{}
	case "truck":
		return Truck{}
	case "bike":
		return Bike{}
	default:
		return nil
	}
}
