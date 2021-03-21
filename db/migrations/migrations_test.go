package migrations_test

import (
	"testing"

	"github.com/nstoker/gorocktrack/internal/pkg/localenv"
	"github.com/nstoker/gorocktrack/internal/pkg/migrations"
)

const testEnv = "../../.test.env"

func TestDownUpMigrate(t *testing.T) {
	localenv.SetEnvironment(testEnv)
	dsn, err := localenv.DatabaseDSN()
	if err != nil {
		t.Fatalf("error getting dsn: %s", err)
	}
	if dsn == "" {
		t.Fatalf("did not get DATABASE_URL from environment")
	}

	err = migrations.MigrateDown(dsn)
	if err != nil {
		t.Fatal(err)
	}

	err = migrations.MigrateUp(dsn)
	if err != nil {
		t.Fatal(err)
	}
}
