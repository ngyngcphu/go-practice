package main

import "log"

func main() {
	car := CreateVehicle("car")
	car.Start()
	car.Stop()

	truck := CreateVehicle("truck")
	truck.Start()
	truck.Stop()

	bike := CreateVehicle("bike")
	bike.Start()
	bike.Stop()

	unknown := CreateVehicle("airplane")
	if unknown == nil {
		log.Println("Error: Unknown vehicle type")
	}
}
