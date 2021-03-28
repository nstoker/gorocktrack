package config

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/nstoker/gorocktrack/app"
	"github.com/sirupsen/logrus"
)

// var DB *sql.DB

// InitDatabase initialises the database for use
func InitDatabase() {
	var err error
	dsn := databaseURI()
	app.DB, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	if err = app.DB.Ping(); err != nil {
		logrus.Errorf("databaseURI %+v", dsn)
		panic(err)
	}

	if err = MigrateUp(dsn); err != nil {
		logrus.Fatalf("migration failure: %v", err)
	}

	seed()
}
