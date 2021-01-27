package utils

import (
	"fiapi/models"

	"github.com/gofiber/fiber/v2"
)

// ErrorHandler is used to catch error thrown inside the routes by ctx.Next(err)
func ErrorHandler(c *fiber.Ctx, err error) error {
	// Statuscode defaults to 500
	code := fiber.StatusInternalServerError

	// Check if it's an fiber.Error type
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(&models.Response{
		Data:    nil,
		Message: err.Error(),
		Success: false,
	})
}

// JwtError is used for catch any error in Protect() JWT function middleware
func JwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return fiber.NewError(fiber.StatusBadRequest, "Missing or malformed JWT")
	}
	return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired JWT")
}
