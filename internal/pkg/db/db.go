package db

import (
	"database/sql"
	"fmt"

	"github.com/nstoker/gorocktrack/internal/pkg/localenv"
	"github.com/nstoker/gorocktrack/internal/pkg/migrations"
	"github.com/sirupsen/logrus"
)

var (
	dsn string
	db  *sql.DB
)

func init() {
	localenv.SetEnvironment(".env")
	dsn, err := localenv.DatabaseDSN()
	if err != nil {
		logrus.Fatalf("db.init fatal %s", err)
	}
	if dsn == "" {
		logrus.Fatalf("db.init fatal no dsn")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logrus.Fatalf("db.init %s opening database", err)
	}

	err = db.Ping()
	if err != nil {
		logrus.Fatalf("db.init fatal ping %s", err)
	}
}

// Migrate migrates the database.
func Migrate(dsn string) error {
	if dsn == "" {
		return fmt.Errorf("dsn missing")
	}

	err := migrations.MigrateUp(dsn)
	if err != nil {
		logrus.Fatal(err)
	}

	return nil
}
