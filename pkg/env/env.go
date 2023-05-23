package env

import (
	"os"
)

var (
	ServerAddress = Get("SERVER_ADDRESS", ":3000")
)

func Get(key string, defaultValue string) string {

	value, isPresent := os.LookupEnv(key)

	if !isPresent {
		return defaultValue
	}

	return value
}
