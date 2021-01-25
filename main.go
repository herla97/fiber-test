package main

import (
	"fiapi/config"
	"fiapi/db"
	"fiapi/models"
	"fiapi/routers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	db.Connect()
	db.Migrate(&models.User{})

	app := fiber.New(fiber.Config{
		ErrorHandler: config.ErrorHandler,
	})

	app.Use(logger.New())
	app.Use(cors.New())

	// TODO: Create db Context
	// db := database.Connect()
	// app.Use(middlewares.ContextDB(db))

	routers.Setup(app)
	log.Fatal(app.Listen(":3000"))
}
