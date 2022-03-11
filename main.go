package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/williamlim16/backend-trash/database"
	"github.com/williamlim16/backend-trash/routes"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load env file")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":" + port)

}
