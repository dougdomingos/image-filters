package utils

import (
	"fmt"
	"log"
	"os"
)

func GetEnvVar(key string) (string, error) {
	value, isDeclared := os.LookupEnv(key)

	if !isDeclared {
		return "", fmt.Errorf("[ERROR] Required environment variable \"%s\" is not declared", key)
	}

	return value, nil
}

func EnsureRequiredEnvVars(keys []string) error {
	for _, key := range keys {
		if value, err := GetEnvVar(key); err != nil {
			return fmt.Errorf("missing required environment variable: %s", key)
		} else {
			log.Printf("%s: %s", key, value)
		}
	}

	return nil
}
