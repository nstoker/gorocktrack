package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/nstoker/gorocktrack/internal/pkg/controllers"
	"github.com/nstoker/gorocktrack/internal/pkg/db"
	"github.com/nstoker/gorocktrack/internal/pkg/localenv"
)

func checkenv() (string, string) {
	port := os.Getenv("PORT")
	if port == "" {
		logrus.Fatal("Environment variable PORT missing")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		logrus.Fatal("Environment variable DATABASE_URL misssing")
	}

	if os.Getenv("ADMIN_EMAIL") == "" {
		logrus.Fatal("Environment variable ADMIN_EMAIL missing")
	}
	if os.Getenv("ADMIN_NAME") == "" {
		logrus.Fatal("Environment variable ADMIN_NAME missing")
	}
	if os.Getenv("ADMIN_PASS") == "" {
		logrus.Fatal("Environment variable ADMIN_PASS missing")
	}

	return dsn, port
}

func main() {
	logrus.Info("Checking environment")
	localenv.SetEnvironment(".env")

	logrus.Info("Initialising")

	dsn, port := checkenv()

	err := db.Migrate(dsn)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("Starting up")

	s := controllers.Server{}

	err = s.InitialiseDatabase(dsn)
	if err != nil {
		logrus.Error(err)
	}

	err = s.InitialiseRouter()
	if err != nil {
		logrus.Error(err)
	}

	logrus.Fatalln(s.Run(port))
}
