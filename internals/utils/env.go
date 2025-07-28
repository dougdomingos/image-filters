package utils

import (
	"fmt"
	"os"
)

func GetEnvVar(key string) (string, error) {
	value, isDeclared := os.LookupEnv(key)

	if !isDeclared {
		return "", fmt.Errorf("[ERROR] Required environment variable \"%s\" is not declared", key)
	}

	return value, nil
}