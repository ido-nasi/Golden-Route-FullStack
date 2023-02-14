package initializers

import (
	"log"
	"os"
	"server/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Makes initial connection to the db
func ConnectToDB() {
	var err error

	// Connects using the dsn in the .env file
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Couldn't connect to database")
	}
}

// Migrates the requested structs into a table in the Database
func MigrateModels() {
	err := DB.AutoMigrate(&models.Flight{})
	if err != nil {
		log.Fatal("Couldn't connect to database")
	}
}
