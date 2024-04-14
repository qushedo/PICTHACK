package main

import (
	"backend/internal/database"
	"backend/internal/routes"
	"github.com/gofiber/fiber/v3"
	"os"
)

func main() {
	database.Connect()

	port := os.Getenv("BACKEND_PORT")
	app := fiber.New()
	routes.Setup(app)

	if err := app.Listen(":" + port); err != nil {
		panic(err)
	}
}
