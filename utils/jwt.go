package utils

import (
	"fiapi/config"

	"github.com/dgrijalva/jwt-go"
)

// TokenGenerator Function for token generation with jwt.
// Receives as parameters the standard claims struct.
// See more: https://tools.ietf.org/html/rfc7519#section-4.1
func TokenGenerator(claims jwt.StandardClaims) (string, error) {
	// Generate token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Get the complete, signed token.
	return token.SignedString([]byte(config.Env("JWT_SECRET")))
}
