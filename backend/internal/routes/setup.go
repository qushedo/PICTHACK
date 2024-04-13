package routes

import (
	"backend/internal/api"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func Setup(app *fiber.App) {
	app.Use(cors.New())
	apiGroup := app.Group("/api", cors.New())

	apiGroup.Get("/ping", api.Ping)
}
