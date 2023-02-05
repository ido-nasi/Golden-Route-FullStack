package main

import (
	"os"
	"server/controllers"
	"server/initializers"

	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.MigrateModels()
}

func main() {
	app := fiber.New()

	app.Get("/", controllers.HomePage)
	app.Get("/allFlights", controllers.GetAllFlights)
	app.Post("/calculate", controllers.Calculate)

	app.Listen(":" + os.Getenv("PORT"))
}
