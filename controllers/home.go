package controllers

import (
	"fiapi/models"

	"github.com/gofiber/fiber/v2"
)

// Hello hanlde api status
func Hello(c *fiber.Ctx) error {
	return c.JSON(&models.Response{
		Data:    nil,
		Message: "Hello i'm ok!",
		Success: true,
	})
}
