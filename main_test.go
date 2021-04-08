package main_test

import (
	"os"
	"testing"

	"github.com/nstoker/gorocktrack/app"
	"github.com/nstoker/gorocktrack/config"
	"github.com/sirupsen/logrus"
)

func TestMain(m *testing.M) {
	config.SetEnvironmentFile(".test.env")
	config.InitDatabase()
	defer app.DB.Close()

	if app.DB == nil {
		logrus.Fatal("no database connection")
	}

	code := m.Run()

	os.Exit(code)
}
