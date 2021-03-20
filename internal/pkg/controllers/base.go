package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

// Server holds the database and router settings
type Server struct {
	DB          *sqlx.DB
	DatabaseURI string
	Router      *mux.Router
}

// InitialiseDatabase initialises the database
func (s *Server) InitialiseDatabase(databaseURI string) error {
	var err error

	db, err := sqlx.Open("postgres", databaseURI)
	if err != nil {
		return fmt.Errorf("server initialise %w", err)
	}
	if err := db.Ping(); err != nil {
		return fmt.Errorf("server pinging %w", err)
	}

	// Run any database migrations needed

	s.DB = db
	s.DatabaseURI = databaseURI

	return nil
}

// InitialiseRouter intialises the router
func (s *Server) InitialiseRouter() error {
	s.Router = mux.NewRouter()

	s.initialiseRoutes()

	return nil
}

// Run server, run. See server run
func (s *Server) Run(addr string) error {
	if addr == "" {
		return fmt.Errorf("missing port")
	}

	logrus.Infof("Listening on %s\n", addr)

	err := http.ListenAndServe(":"+addr, s.Router)
	return fmt.Errorf(err.Error())
}
