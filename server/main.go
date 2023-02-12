package main

import (
	"os"
	"server/controllers"
	"server/initializers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.MigrateModels()
}

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("FRONTEND_URL"),
		AllowMethods: "GET,POST,PUT,OPTIONS",
		AllowHeaders: "Origin,Content-Type,Accept,Access-Control-Allow-Origin",
	}))

	app.Get("/api/v1/allFlights", controllers.GetAllFlights)
	app.Post("/api/v1/calculate", controllers.Calculate)
	app.Post("api/v1/weatherApi", controllers.WeatherApiController)

	app.Listen(":" + os.Getenv("PORT"))
}
