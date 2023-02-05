package initializers

import (
	"log"
	"os"
	"server/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Couldn't connect to database")
	}
}

func MigrateModels() {
	err := DB.AutoMigrate(&models.Flight{})
	if err != nil {
		log.Fatal("Couldn't connect to database")
	}
}
