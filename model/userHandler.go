package model

import (
	"net/http"

	"github.com/nstoker/gorocktrack/config"
)

func UserIndex(w http.ResponseWriter, r *http.Request) {
	users, err := ShowAllUsers()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	config.TPL.ExecuteTemplate(w, "users.index.gohtml", users)
}
