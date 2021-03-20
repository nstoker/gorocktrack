package seeders

import (
	"fmt"

	"github.com/nstoker/gorocktrack/internal/pkg/localenv"
)

func userSeed(dsn string) error {
	if dsn == "" {
		return fmt.Errorf("dsn missing")
	}

	envVars := localenv.RequiredEnvironmentVariables()
	admin, err := localenv.GetEnvironmentValues(envVars)
	if err != nil {
		return err
	}
	if admin == nil {
		return fmt.Errorf("admin environment variable results missing")
	}

	return fmt.Errorf("not yet implemented")
}
