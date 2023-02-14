package controllers

import (
	"fmt"
	"net/http"
	"server/initializers"
	"server/models"
	"server/physics"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Handles Post request to the /calculate endpoint
// Implements the physics calculator logic.
func Calculate(c *fiber.Ctx) error {
	mp := MassParam{}
	err := c.BodyParser(&mp)

	// Parsing parameters from request body
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	// validating mass value & converting to float64 type
	mass := mp.validateAndConvertMass()

	// Handles invalid mass value
	if mass < 0 {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Invalid mass value"})
		return err
	}

	// Calculating the needed metrics & packing into a struct
	excessMass, time := physics.TakeOffTime(mass)
	distance := physics.TakeOffDistance(mass)
	flight := models.NewFlight(mass, distance, excessMass, time)

	// Inserting record to the DB
	err = initializers.DB.Create(&flight).Error

	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "couldn't create flight record"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{"flightData": flight})
	return nil
}

// for own testing purposes
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

// Handles post request to the /weatherApi endpoint
func WeatherApiController(ctx *fiber.Ctx) error {
	// Parse request body to struct
	var requestParams RequestParams
	err := ctx.BodyParser(&requestParams)

	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}
	// Extract date with format YYYY-MM-DD
	extractedDate := strings.Split(requestParams.Date, "T")[0]

	// Makes a request to the remote weather api
	getString := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&hourly=temperature_2m&timezone=auto&start_date=%s&end_date=%s", requestParams.Latitude, requestParams.Longitude, extractedDate, extractedDate)
	response, err := http.Get(getString)
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "couldn't contact the api"})
		return err
	}

	// Parsing the apiResponse body to a designated struct
	var apiResponse WeatherApiResponse
	apiResponse.UnmarshalResponse(response)

	// Checking for error returned from the api
	if apiResponse.Error == true {
		ctx.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": apiResponse.Reason})
		return err
	}

	// Logic Block - finds the hours where the temperature enables take off
	// Stroing them in a slice.
	var takeOffValidHours []string
	for i := range apiResponse.Hourly.Time {
		time := strings.Split(apiResponse.Hourly.Time[i], "T")
		temprature := apiResponse.Hourly.Temperature[i]

		if float64(15) <= temprature && temprature <= float64(30) {
			takeOffValidHours = append(takeOffValidHours, time[1])
		}
	}

	// In case there are valid take off hours
	if len(takeOffValidHours) > 0 {
		ctx.Status(http.StatusOK).JSON(&fiber.Map{"hours": takeOffValidHours})
		return nil
	}

	// If there aren't any valid take off hours, calculate and return the average temperature
	sum := 0.0
	for _, temp := range apiResponse.Hourly.Temperature {
		sum += temp
	}
	sum /= 24

	ctx.Status(http.StatusOK).JSON(&fiber.Map{"avgTemp": sum})
	return nil
}
