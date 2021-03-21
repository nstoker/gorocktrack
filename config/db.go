package config

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", databaseURI())
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	logrus.Infoln("Database connection succesful")
}
