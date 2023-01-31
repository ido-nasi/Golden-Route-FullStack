package main

import (
	"fmt"

	"github.com/ido-nasi/GoldenLane/initializers"
	"github.com/ido-nasi/GoldenLane/physics"
)

func init() {
	initializers.LoadEnvVariables()
	// initializers.ConnectToDB()
}

func main() {
	// router := gin.Default()

	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(200, "Hello world")
	// })

	fmt.Println(physics.FlightDistance(100_000))

	// router.Run()
}
