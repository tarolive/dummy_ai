package env

import (
	"os"
)

var (
	ServerAddress = _get("SERVER_ADDRESS", ":3000")
)

func _get(key string, defaultValue string) string {

	if value, isPresent := os.LookupEnv(key); isPresent {

		return value
	}

	return defaultValue
}
