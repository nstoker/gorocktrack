package model_test

import (
	"testing"

	"github.com/nstoker/gorocktrack/app"
	"github.com/nstoker/gorocktrack/model"
)

func TestAddAdminUser(t *testing.T) {
	adminEmail := "test_admin@example.com"
	adminName := "An Admin"
	adminPassword := "changeM3!"
	if app.DB == nil {
		t.Fatal("database not connected")
	}
	u, err := model.AddAdminUser(adminName, adminEmail, adminPassword, true)
	if err != nil {
		t.Fatal(err)
	}

	if u.Name != adminName {
		t.Errorf("expected %s, got '%s'", adminName, u.Name)
	}

	if u.Email != adminEmail {
		t.Errorf("expected %s, got '%s'", adminEmail, u.Email)
	}

	isAdmin, err := u.Admin()
	if err != nil {
		t.Error(err)
	}
	if !isAdmin {
		t.Errorf("expected isAdmin to be true, got %v", isAdmin)
	}
}

func TestGetAllUsers(t *testing.T) {
	users, err := model.GetAllUsers()
	if err != nil {
		t.Error(err)
	}
	if len(users) == 0 {
		t.Error("Expected user records, nothing returned")
	}
}
