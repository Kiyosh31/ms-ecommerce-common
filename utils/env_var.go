package utils

import (
	"errors"
	"os"
)

func GetEnvVar(envVar string) (string, error) {
	value := os.Getenv(envVar)
	if value == "" {
		err := errors.New("Missing env var: " + envVar)
		return "", err
	}

	return value, nil
}
