package physics

import (
	"log"
	"math"
)

var engineForce float64 = 100_000
var baseMass float64 = 35_000 // place + team
var maxTakeOffTime float64 = 60
var flightSpeed float64 = 140
var maxCargoMassForTimeLimit float64 = 7857.142857 // The calculation in the README.md

func Acceleration(cargoMass float64) float64 {

	if cargoMass < 0 {
		log.Println("cargoMass argument MUST be positive.")
		return -1
	}

	return engineForce / (cargoMass + baseMass)
}

func TakeOffTime(cargoMass float64) (float64, float64) {
	/*
		ARGS:
			- first arg: the mass to reduce in case we are over the max time limit
			- second arg: the take-off time (in seconds)
	*/
	if cargoMass < 0 {
		log.Println("cargoMass argument MUST be positive.")
		return -1, -1
	}

	a := Acceleration(cargoMass)
	t := flightSpeed / a

	if t > maxTakeOffTime {
		return cargoMass - maxCargoMassForTimeLimit, t
	}

	return 0, t
}

func TakeOffDistance(cargoMass float64) float64 {

	if cargoMass < 0 {
		log.Println("cargoMass argument MUST be positive.")
		return -1
	}

	a := Acceleration(cargoMass)
	_, t := TakeOffTime(cargoMass)

	return 0.5 * a * math.Pow(t, 2)
}
