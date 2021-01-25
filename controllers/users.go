package controllers

import (
	"fiapi/db"
	"fiapi/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func SignInUser(c *fiber.Ctx) error {
	// User Model
	user := &models.User{}

	db := db.DB
	if err := c.BodyParser(user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Review your input")
	}

	password := user.Password

	if err := db.Where(&models.User{Username: user.Username}).Find(user).Error; err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid email or password")
	}

	// confirm that the given password matches the hashed password from the db
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid email or password")
	}

	cc := &models.CustomClaims{
		ID: user.ID,
		// Role: "artist",
		// Type: "session",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 720).Unix(),
		},
	}

	tokenString, err := cc.TokenGenerator()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusAccepted).JSON(&models.Response{
		Data:    tokenString,
		Message: "User Authorized",
		Success: true,
	})
}

// CreateUser new user
func CreateUser(c *fiber.Ctx) error {
	// User Model
	user := &models.User{}

	db := db.DB
	if err := c.BodyParser(user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Review your input")
	}

	if err := user.HashPassword(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Couldn't hash password")
	}

	if err := db.Create(&user).Error; err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Couldn't create user")
	}

	user.Password = ""

	return c.Status(fiber.StatusCreated).JSON(&models.Response{
		Data:    user,
		Message: "User Created",
		Success: true,
	})
}
