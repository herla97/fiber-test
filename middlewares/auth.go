package middlewares

import (
	"fiapi/config"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.Env("JWT_SECRET")),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return fiber.NewError(fiber.StatusBadRequest, "Missing or malformed JWT")
		// return c.Status(fiber.StatusBadRequest).
		// 	JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}

	return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired JWT")
	// JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

// func ContextDB(db *gorm.DB) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		c.Set("db", )
// 		return c.Next()
// 	}
// }
