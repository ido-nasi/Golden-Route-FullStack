package controllers

import (
	"regexp"
	"strconv"
)

// Struct to extract mass from request body
type MassParam struct {
	Mass string `json:"mass"`
}

// Struct function to validate and convert the extracted mass value
// In case of invalid mass value (string or invalid float format), return -1
func (m *MassParam) validateAndConvertMass() float64 {
	re := regexp.MustCompile("^[+\\-]?(?:(?:0|[1-9]\\d*)(?:\\.\\d*)?|\\.\\d+)$")
	mass := re.FindString(m.Mass)

	if mass == "" {
		return -1
	}

	value, err := strconv.ParseFloat(mass, 64)

	if err != nil {
		return -1
	}

	return value
}
