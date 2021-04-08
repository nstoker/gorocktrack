package model_test

import (
	"os"
	"testing"

	"github.com/nstoker/gorocktrack/config"
)

func TestMain(m *testing.M) {
	config.SetEnvironmentFile("../.test.env")
	config.InitDatabase()

	code := m.Run()

	os.Exit(code)
}
