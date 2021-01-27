package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// VerifyPassword compares a hashed password with plaintext password
func VerifyPassword(hash string, raw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw))
}
