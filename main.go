package main

import (
	"fiapi/config"
	"fiapi/config/db"
	"fiapi/controllers"
	"fiapi/models"
	"fiapi/routers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db.Connect()
	db.Migrate(&models.User{})

	app := fiber.New(fiber.Config{
		// TODO: Agregar más configuraciones
		ErrorHandler: controllers.ErrorHandler,
	})

	app.Use(cors.New())

	routers.Setup(app)
	log.Fatal(app.Listen(":" + config.Env("PORT")))
}
