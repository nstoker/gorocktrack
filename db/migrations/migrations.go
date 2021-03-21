package migrations

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/nstoker/gorocktrack/internal/pkg/localenv"

	// These files are needed for the database and migration utilities
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func prepareMigration(dsn string) (*migrate.Migrate, error) {
	if dsn == "" {
		return nil, fmt.Errorf("dsn missing")
	}

	mp := fmt.Sprintf("file://%s/migrations/files/", localenv.RootDir())

	return migrate.New(mp, dsn)
}

// MigrateDown will reverse the migrations
func MigrateDown(dsn string) error {
	m, err := prepareMigration(dsn)
	if err != nil {
		return fmt.Errorf("migrate.down prepare: %s", err)
	}

	if err = m.Down(); err != nil {
		if err != migrate.ErrNoChange {
			return fmt.Errorf("performing migrate.down: %s", err)
		}
	}

	return nil
}

// MigrateUp will migrate the database
func MigrateUp(dsn string) error {
	m, err := prepareMigration(dsn)
	if err != nil {
		return fmt.Errorf("migrate.up prepare error: %s", err)
	}

	if err = m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			return fmt.Errorf("error performing migrate.up: %s", err)
		}
	}

	return nil
}
