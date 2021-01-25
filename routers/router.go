package routers

import (
	"fiapi/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Setup Setup Routes
func Setup(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", controllers.Hello)

	api.Post("/sign_up", controllers.CreateUser)
	api.Post("/sign_in", controllers.SignInUser)
}
