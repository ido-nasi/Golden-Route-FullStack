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
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/", controllers.HomePage)
	app.Get("/allFlights", controllers.GetAllFlights)
	app.Post("/calculate", controllers.Calculate)

	app.Listen(":" + os.Getenv("PORT"))
}
