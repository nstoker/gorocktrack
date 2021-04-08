package config

import (
	"database/sql"
	"fmt"

	"github.com/nstoker/gorocktrack/model"
	"github.com/sirupsen/logrus"
)

// SeedDatabase populates database with key values
func SeedDatabase() {
	seed_admin()
}

func seed_admin() {
	admin, err := GetRequiredEnvironmentVariables([]string{"ADMIN_EMAIL", "ADMIN_NAME", "ADMIN_PASS"})
	if err != nil {
		logrus.Panicf(err.Error())
	}

	user, err := model.FindByEmail(admin["ADMIN_EMAIL"])

	switch err {
	case sql.ErrNoRows:
		addAdminRecord(admin)
	case nil:
		checkAdmin(user)
	default:
		logrus.Fatalf(err.Error())
	}
}

func checkAdmin(user *model.User) {
	isAdmin, err := user.Admin()
	if err != nil {
		logrus.Panic(err)
	}
	if !isAdmin {
		logrus.Panic(fmt.Errorf("user record exists but is not admin"))
	}
}

func addAdminRecord(admin map[string]string) {
	user, err := model.AddAdminUser(admin["ADMIN_NAME"], admin["ADMIN_EMAIL"], admin["ADMIN_PASS"], true)
	if err != nil {
		logrus.Panic(err)
	}

	checkAdmin(user)
}
