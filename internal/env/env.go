package env

import (
	"os"
	"strconv"
)

func GetEnv(key, defaultValue string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		return defaultValue
	}

	return val
}

func GetEnvAsInt(key string, defaultValue int) int {
	val, ok := os.LookupEnv(key)

	if !ok {
		return defaultValue
	}

	valInt, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue // Return default value if conversion fails
	}

	return valInt
}
