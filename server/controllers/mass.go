package controllers

import (
	"regexp"
	"strconv"
)

// struct to extract mass from request body
type MassParam struct {
	Mass string `json:"mass"`
}

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
