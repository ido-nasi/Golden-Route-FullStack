package physics

import (
	"log"
	"math"
)

var engineForce float64 = 100_000
var baseMass float64 = 35_000
var maxTakeOffTime float64 = 60
var flightSpeed float64 = 140
var maxMassForTimeLimit float64 = 7857.142857 // The calculation in the README.md

func Acceleration(cargoMass float64) float64 {

	if cargoMass < 0 {
		log.Fatal("cargoMass argument MUST be positive.")
		return -1
	}

	return engineForce / (cargoMass + baseMass)
}

func FlightTime(cargoMass float64) (int, float64) {
	a := Acceleration(cargoMass)

	if a == -1 {
		log.Fatal("cargoMass argument MUST be positive.")
		return -1, -1
	}

	t := flightSpeed / a

	if t > maxTakeOffTime {
		return -1, (cargoMass - maxMassForTimeLimit)
	}

	return 1, t
}

func FlightDistance(cargoMass float64) float64 {
	a := Acceleration(cargoMass)
	status, t := FlightTime(cargoMass)

	if status == -1 && t == -1 {
		log.Fatal("cargoMass argument MUST be positive.")
		return -1
	}

	if status == -1 && t != -1 {
		log.Printf("Mass to lose %f", t)
		return t
	}

	return 0.5 * a * math.Pow(t, 2)
}
