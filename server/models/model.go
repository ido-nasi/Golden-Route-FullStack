package models

// Flight struct to represent the Flight table in the Database
type Flight struct {
	FlightId        uint    `gorm:"primary key;autoIncrement;not null" json:"id"`
	CargoMass       float64 `gorm:"not null" json:"mass"`
	TakeOffDistance float64 `gorm:"not null" json:"distance"`
	ExcessCargoMass float64 `gorm:"not null" json:"excessMass"`
	TakeOffTime     float64 `gorm:"not null" json:"time"`
}

// Constructor to create a Flight instance without the FlightId
// FlightId is a bigserial (sequence) created automatically by the database
func NewFlight(cargoMass float64, takeOffDistance float64, excessCargoMass float64, takeOffTime float64) Flight {
	return Flight{CargoMass: cargoMass, TakeOffDistance: takeOffDistance, ExcessCargoMass: excessCargoMass, TakeOffTime: takeOffTime}
}
