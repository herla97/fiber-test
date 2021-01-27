package main

import (
	"fiapi/config"
	"fiapi/config/db"
	"fiapi/models"
	"fiapi/routers"
	"fiapi/utils"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db.Connect()
	db.Migrate(&models.User{})

	app := fiber.New(fiber.Config{
		// TODO: Agregar más configuraciones
		ErrorHandler: utils.ErrorHandler,
	})

	app.Use(cors.New())

	routers.Setup(app)
	log.Fatal(app.Listen(":" + config.Env("PORT")))
}
