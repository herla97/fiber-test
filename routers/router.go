package routers

import (
	"fiapi/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", controllers.Hello)
}
