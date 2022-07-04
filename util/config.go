package util

import (
	"os"

	"github.com/joho/godotenv"
)

// use godot package to load/read the .env file and
// return the value of the key
func LoadEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load("app.env")
	if err != nil {
		return ""
	}
	return os.Getenv(key)
}
