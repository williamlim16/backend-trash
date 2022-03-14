package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/williamlim16/backend-trash/models"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("dsn")
	db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if error != nil {
		log.Fatal("Could not connect tot he database")
	} else {
		log.Output(1, "Connection success!")
		log.Println(db)
	}
	db.AutoMigrate(
		&models.User{},
		&models.TrashCan{},
		&models.Trash{},
	)
	DB = db
}
