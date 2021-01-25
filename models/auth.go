package models

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// CustomClaims is used by TokenGenerator Func
// See Documentation: https://godoc.org/github.com/dgrijalva/jwt-go
type CustomClaims struct {
	ID uuid.UUID `json:"id"`
	// Role string    `json:"role" rw:"r"` // TODO: Review
	// Type string    `json:"type" rw:"r"` // TODO: Review
	jwt.StandardClaims
}

// TokenGenerator token generation function.
func (c *CustomClaims) TokenGenerator() (string, error) {
	// Generate token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	// Get the complete, signed token.
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
