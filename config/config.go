package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Env func to get env value
func Env(key string) string {
	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
