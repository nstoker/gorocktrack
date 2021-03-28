package config

import (

	// These files are needed for the database and migration utilities
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func prepareMigration(dsn string) (*migrate.Migrate, error) {
	if dsn == "" {
		return nil, fmt.Errorf("DSN missing")
	}

	mp := fmt.Sprintf("file://%s/migrations/files", RootDirectory())

	return migrate.New(mp, dsn)
}

// MigrateDown will reverse a migration
func MigrateDown(dsn string) error {
	m, err := prepareMigration(dsn)
	if err != nil {
		return fmt.Errorf("migrate.down prepare: %v", err)
	}

	if err = m.Down(); err != nil {
		if err != migrate.ErrNoChange {
			return fmt.Errorf("performing migrate.down: %v", err)
		}
	}

	return nil
}

// MigrateUp will apply all migrations needed
func MigrateUp(dsn string) error {
	m, err := prepareMigration(dsn)
	if err != nil {
		return fmt.Errorf("migrate.up prepare: %v", err)
	}

	if err = m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			return fmt.Errorf("performing migrate.up: %v", err)
		}
	}

	return nil
}
