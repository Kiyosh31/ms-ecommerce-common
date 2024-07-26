package utils

import (
	"errors"
	"os"
)

func GetEnvVar(envVar string) (string, error) {
	value := os.Getenv(envVar)
	if value == "" {
		err := errors.New("You must provide" + envVar + "in environment variables")
		return "", err
	}

	return value, nil
}
