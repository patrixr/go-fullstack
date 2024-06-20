package core

import (
	"fmt"
	"os"
)

// ReadEnv retrieves the value of the environment variable named by the key.
// If the environment variable is not set, or is empty, the defaultValue is returned.
func ReadEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return defaultValue
}

// ReadEnvStrict retrieves the value of the environment variable named by the key.
// It panics if the environment variable is not set or is empty.
func ReadEnvStrict(key string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	panic(fmt.Sprintf("Environment variable %s is required", key))
}
