package config

import (
	"os"
	"strconv"
)

func GetEnvAsFloat64(key string, fallback float32) float32 {
	valueStr := os.Getenv(key)
	if value, err := strconv.ParseFloat(valueStr, 32); err == nil {
		return float32(value)
	}
	return fallback
}

func GetEnvAsBool(key string, fallback bool) bool {
	valueStr := os.Getenv(key)
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return fallback
}

func GetEnvAsString(key string) string {
	return os.Getenv(key)
}
