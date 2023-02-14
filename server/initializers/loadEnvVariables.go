package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// Load environment variables from the .env file
func LoadEnvVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
