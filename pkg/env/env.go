package env

import (
	"os"
)

var (
	ServerAddress = GetString("SERVER_ADDRESS", ":3000")
)

func GetString(key string, defaultValue string) string {

	value := os.Getenv(key)

	if value == "" {
		value = defaultValue
	}

	return value
}
