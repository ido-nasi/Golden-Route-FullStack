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

func WeatherApiController(ctx *fiber.Ctx) error {
	// pasre request params
	var requestParams RequestParams
	err := ctx.BodyParser(&requestParams)

	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}
	extractedDate := strings.Split(requestParams.Date, "T")[0]

	// request the remote api
	getString := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&hourly=temperature_2m&timezone=auto&start_date=%s&end_date=%s", requestParams.Latitude, requestParams.Longitude, extractedDate, extractedDate)
	response, err := http.Get(getString)
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "couldn't contact the api"})
		return err
	}

	var apiResponse WeatherApiResponse
	apiResponse.UnmarshalResponse(response)

	if apiResponse.Error == true {
		ctx.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": apiResponse.Reason})
		return err
	}

	// logic - finds the hours where the temperature is sufficient to take off
	var takeOffValidHours []string
	for i := range apiResponse.Hourly.Time {
		time := strings.Split(apiResponse.Hourly.Time[i], "T")
		temprature := apiResponse.Hourly.Temperature[i]

		if float64(15) <= temprature && temprature <= float64(30) {
			takeOffValidHours = append(takeOffValidHours, time[1])
		}
	}

	if len(takeOffValidHours) > 0 {
		ctx.Status(http.StatusOK).JSON(&fiber.Map{"hours": takeOffValidHours})
		return nil
	}

	sum := 0.0
	for _, temp := range apiResponse.Hourly.Temperature {
		sum += temp
	}
	sum /= 24

	ctx.Status(http.StatusOK).JSON(&fiber.Map{"avgTemp": sum})
	return nil
}
