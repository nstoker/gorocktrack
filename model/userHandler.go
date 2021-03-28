package model

import (
	"net/http"

	"github.com/nstoker/gorocktrack/app"
	"github.com/sirupsen/logrus"
)

func UserIndex(w http.ResponseWriter, r *http.Request) {
	users, err := ShowAllUsers()
	if err != nil {
		logrus.Errorf("UserIndex: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	app.TPL.ExecuteTemplate(w, "users_index.gohtml", users)
}
