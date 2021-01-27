package controllers

import (
	"fiapi/config/db"
	"fiapi/models"
	"fiapi/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// AuthUsers default implementation.
func AuthUsers(c *fiber.Ctx) error {
	auth := &models.Auth{} // Auth Model
	user := &models.User{} // User Model

	db := db.Connection // Instance of connection to db
	if err := c.BodyParser(auth); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "review your input")
	}

	// User search
	if err := db.Where("email = ?", auth.Username).Find(user).Error; err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "email o contraseña invalida")
	}

	matches, err := utils.VerifyPassword(auth.Password, user.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if !matches {
		return fiber.NewError(fiber.StatusUnauthorized, "email o contraseña invalida")
	}

	// Token generator. Return token string
	token, err := utils.TokenGenerator(jwt.StandardClaims{
		Id:        user.ID.String(),
		Audience:  "user", // TODO: Agregar manejo de roles con Casbin
		ExpiresAt: time.Now().Add(time.Hour * 720).Unix(),
	})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusAccepted).JSON(&models.Response{
		Data:    token,
		Message: "User Authorized",
		Success: true,
	})
}
