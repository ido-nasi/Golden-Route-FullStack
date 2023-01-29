package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ido-nasi/GoldenLane/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello world")
	})

	router.Run()
}
