package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func setEnvironmentFile(envFile string) error {
	if _, err := os.Stat(envFile); err == nil {
		godotenv.Load(envFile)
	}

	return checkRequiredEnvVars()
}

func checkRequiredEnvVars() error {
	expectedVariables := []string{"DATABASE_URL"}
	missingVariables := []string{}

	for _, envVar := range expectedVariables {
		value := os.Getenv(envVar)
		if value == "" {
			missingVariables = append(missingVariables, envVar)
		}
	}

	if len(missingVariables) > 0 {
		return fmt.Errorf("missing from environment: %s", strings.Join(missingVariables, ", "))
	}

	return nil
}

func databaseURI() string {
	return os.Getenv("DATABASE_URL")
}
