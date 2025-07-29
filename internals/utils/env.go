package utils

import (
	"fmt"
	"log"
	"os"
)

// EnsureRequiredEnvVars verifies if the variables specified through the
// input slice are declared in the environment. If any variable is not
// declared, returns an error indicating the missing variable.
func EnsureRequiredEnvVars(keys []string) error {
	for _, key := range keys {
		if value, isDeclared := os.LookupEnv(key); !isDeclared {
			return fmt.Errorf("missing required environment variable: %s", key)
		} else {
			log.Printf("[ENV] %s: %s", key, value)
		}
	}

	return nil
}
