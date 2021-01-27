package routers

import (
	"fiapi/controllers"
	"fiapi/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Setup Setup Routes
func Setup(app *fiber.App) {
	// Logger Middleware all App
	app.Use(logger.New())

	// If you need add logger config
	// app.Use(logger.New(logger.Config{
	// 	Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
	// 	TimeFormat: "02-Jan-2006",
	// 	TimeZone:   "America/Mexico_City",
	// }))

	// App Grouping
	// If you need add logger by group
	// api := app.Group("/api", logger.New()) // Add logger by group

	api := app.Group("/api")
	// api.Get("/", controllers.Hello)

	api.Post("/u/sign_in", controllers.AuthUsers)
	api.Post("/u/sign_up", controllers.UsersCreate)

	user := app.Group("/api/u").Use(middlewares.Protected())
	user.Get("/", controllers.Hello)
}
