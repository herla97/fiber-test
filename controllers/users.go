package controllers

import (
	"fiapi/config/db"
	"fiapi/models"

	"github.com/gofiber/fiber/v2"
)

// UsersCreate new user
func UsersCreate(c *fiber.Ctx) error {
	user := &models.User{} // User Model
	db := db.Connection    // Instance of connection to db

	if err := c.BodyParser(user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Review your input")
	}

	if err := db.Create(user).Error; err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&models.Response{
		Data:    user,
		Message: "User Created",
		Success: true,
	})
}
