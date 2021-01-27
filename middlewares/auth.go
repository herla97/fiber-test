package middlewares

import (
	"fiapi/config"
	"fiapi/controllers"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.Env("JWT_SECRET")),
		ErrorHandler: controllers.JwtError,
	})
}
