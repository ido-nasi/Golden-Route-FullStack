package controllers

import (
	"net/http"
	"regexp"
	"server/initializers"
	"server/models"
	"server/physics"
	"strconv"

	"github.com/gofiber/fiber/v2"
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

func HomePage(c *fiber.Ctx) error {
	return c.JSON("hello world")

}

func Calculate(c *fiber.Ctx) error {
	mp := MassParam{}
	err := c.BodyParser(&mp)

	// parsing error from request body
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	// validating value & converting to float64 type
	mass := mp.validateAndConvertMass()

	if mass < 0 {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Invalid mass value"})
		return err
	}

	excessMass, time := physics.TakeOffTime(mass)
	distance := physics.TakeOffDistance(mass)

	flight := models.NewFlight(mass, distance, excessMass, time)
	err = initializers.DB.Create(&flight).Error

	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "couldn't create flight record"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{"flightData": flight})
	return nil
}

func GetAllFlights(c *fiber.Ctx) error {

	flights := []models.Flight{}

	err := initializers.DB.Find(&flights).Error

	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "couldn't retrieve records"})
		return err
	}

	c.Status(http.StatusOK).JSON(flights)

	return nil
}
