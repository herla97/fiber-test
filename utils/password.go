package utils

import (
	"github.com/matthewhartstonge/argon2"
)

// HashPassword Function for password generation. Using Argon2
// See more: https://github.com/matthewhartstonge/argon2
// https://pkg.go.dev/golang.org/x/crypto/argon2
func HashPassword(password string) (string, error) {
	argon := argon2.DefaultConfig() // Argon2 default config

	// Argon2 encode password
	encoded, err := argon.HashEncoded([]byte(password))
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

// VerifyPassword compares a hashed password with plaintext password
func VerifyPassword(raw string, hash string) (bool, error) {
	// return bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw))
	ok, err := argon2.VerifyEncoded([]byte(raw), []byte(hash))
	if err != nil {
		return ok, err
	}
	return ok, err
}
