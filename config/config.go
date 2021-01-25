package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "dev".
var ENV = os.Setenv("GO_ENV", "dev")

// func init() {
// 	// load .env file
// 	if err := godotenv.Load(".env"); err != nil {
// 		fmt.Print("Error loading .env file")
// 	}

// 	// Print App Enviroment
// 	fmt.Printf("ENV: %v \n", ENV)
// }

// Env func to get env value
func Env(key string) string {
	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
