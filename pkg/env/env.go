package env

import (
	"os"
)

var (
	ServerAddress = Get("SERVER_ADDRESS", ":3000")
)

func Get(key string, defaultValue string) string {

	if value, isPresent := os.LookupEnv(key); isPresent {

		return value
	}

	return defaultValue
}
