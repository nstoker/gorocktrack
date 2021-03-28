package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// SetLogger sets up the logger
func SetLogger() {
	var formatter logrus.Formatter

	if os.Getenv("HOME") == "/home/ns" {
		formatter = &logrus.TextFormatter{}
	} else {
		formatter = &logrus.JSONFormatter{}
	}

	logrus.SetFormatter(formatter)
	logrus.SetReportCaller(true)
}

// RootDirectory will return the app's root directory
func RootDirectory() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))

	return filepath.Dir(d)
}

// GetEnvironmentValues returns an array of selected environment variables
func GetEnvironmentVariables(vars []string) map[string]string {
	results := make(map[string]string)

	for _, environmentVariable := range vars {
		value := os.Getenv(environmentVariable)

		results[environmentVariable] = value
	}

	return results
}

// GetRequiredEnvironmentVariables gets a map of variables, or an error unless all have values
func GetRequiredEnvironmentVariables(vars []string) (map[string]string, error) {
	results := GetEnvironmentVariables(vars)
	errors := []string{}

	for _, envVar := range vars {
		value := results[envVar]
		if value == "" {
			errors = append(errors, envVar)
		}
	}

	if len(errors) > 0 {
		return make(map[string]string), fmt.Errorf("missing from environment: %s", strings.Join(errors, ", "))
	}

	return results, nil
}

// SetEnvironmentFile sets the environment file
func SetEnvironmentFile(envFile string) error {
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
