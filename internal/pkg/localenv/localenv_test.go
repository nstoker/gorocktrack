package localenv_test

import (
	"testing"

	"github.com/nstoker/gorocktrack/internal/pkg/localenv"
)

const testEnvironment = "../../../.test.env"

func TestCheckMissingEnvironment(t *testing.T) {
	localenv.SetEnvironment("")
	dsn, err := localenv.DatabaseDSN()
	if err == nil {
		t.Error("expected to get an error")
	}

	if dsn != "" {
		t.Errorf("did not expect a dsn, got '%s'", dsn)
	}
}

func TestCheckExistingTestEnvironment(t *testing.T) {
	localenv.SetEnvironment(testEnvironment)
	dsn, err := localenv.DatabaseDSN()
	if err != nil {
		t.Fatalf("error: '%v'", err)
	}

	if dsn == "" {
		t.Errorf("did not receive a dsn from environment")
	}
}

func TestCheckExistingDevEnvironment(t *testing.T) {
	localenv.SetEnvironment(testEnvironment)
	dsn, err := localenv.DatabaseDSN()
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	if dsn == "" {
		t.Errorf("did not receive a dsn from environment")
	}
}

func TestGetEnvironmentValues(t *testing.T) {
	localenv.SetEnvironment(testEnvironment)
	expectedEnvironmentVars := localenv.RequiredEnvironmentVariables()

	results, err := localenv.GetEnvironmentValues(expectedEnvironmentVars)
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != len(expectedEnvironmentVars) {
		t.Fatalf("Expected %v result, got %v %v", len(results),
			len(expectedEnvironmentVars), expectedEnvironmentVars)
	}
}
