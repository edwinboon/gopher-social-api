package env

import (
	"os"
	"strconv"
	"time"
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
		return defaultValue
	}

	return valInt
}

func GetEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	valDuration, err := time.ParseDuration(val)
	if err != nil {
		return defaultValue
	}

	return valDuration
}
