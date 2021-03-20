package localenv

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	environment string
	Hash        string
	Block       string
)

// SetEnvironment sets the environment variable for later
func SetEnvironment(envVar string) error {
	environment = envVar

	return godotenv.Load(environment)
}

// DatabaseDSN returns the address to use for the database
func DatabaseDSN() (string, error) {
	err := godotenv.Load(environment)
	if err != nil {
		return "", err
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return "", fmt.Errorf("environment variable 'DATABASE_URL' missing")
	}

	return dsn, nil
}

// GetEnvironmentValues gets an array of environment variables
func GetEnvironmentValues(vars []string) (map[string]string, error) {
	godotenv.Load(environment)
	results := make(map[string]string)
	errors := []string{}

	for _, envVar := range vars {
		value := os.Getenv(envVar)
		if value == "" {
			errors = append(errors, envVar)
		} else {
			results[envVar] = value
		}
	}
	if len(errors) > 0 {

		return nil, fmt.Errorf("missing environment variables: %s", strings.Join(errors, ", "))
	}
	return results, nil
}

func RequiredEnvironmentVariables() []string {
	envVars := []string{"ADMIN_EMAIL", "ADMIN_NAME", "ADMIN_PASS", "HASH", "BLOCK"}
	return envVars
}
