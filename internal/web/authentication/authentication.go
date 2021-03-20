package authentication

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/nstoker/gorocktrack/internal/pkg/models"
	"github.com/sirupsen/logrus"
)

var templates = template.Must(template.ParseGlob("internal/web/authentication/templates/*"))

// Routes sets up the authentication routes
func Routes(r *mux.Router) {
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/login", loginPage).Methods("GET")

	r.HandleFunc("/signup", signupHandler).Methods("POST")
	r.HandleFunc("/signup", signupPage).Methods("GET")

	r.HandleFunc("/logout", logoutHandler).Methods("DESTROY")
}

// loginPage
func loginPage(response http.ResponseWriter, request *http.Request) {
	err := templates.ExecuteTemplate(response, "login.html", nil)
	if err != nil {
		logrus.Infof("internal/web/authentication/loginPage: %v", err)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(fmt.Sprintf("login page handler %s", err)))
		return
	}
}

// signupPage
func signupPage(response http.ResponseWriter, request *http.Request) {
	err := templates.ExecuteTemplate(response, "signup.html", nil)
	if err != nil {
		logrus.Infof("internal/web/authentication/signupPage: %v", err)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(fmt.Sprintf("signup page handler %s", err)))
		return
	}
}

// Login handles user login
func loginHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	pass := request.FormValue("password")
	redirectTarget := "/"
	if name != "" && pass != "" {
		// Check creds
		setSession(name, response)
	}
	http.Redirect(response, request, redirectTarget, http.StatusFound)
}

// Signup to the system
// Eventually the password option will be removed and set after an email has been received.
func signupHandler(w http.ResponseWriter, r *http.Request) {
	creds := &models.Credentials{}
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	ok, err := creds.Signup(name, email, password)
	if !ok {
		logrus.Infof("Error %v", err)
		http.Redirect(w, r, "/", http.StatusBadRequest)
	}

	http.Redirect(w, r, "/", http.StatusCreated)
}

// Logout from the system
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/", http.StatusFound)
}
