package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/nstoker/gorocktrack/internal/pkg/version"
	landing "github.com/nstoker/gorocktrack/internal/web/landing"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		logrus.Fatal("Environment variable PORT missing")
	}

	logrus.Infof("Starting up %s", version.Version)

	r := mux.NewRouter()
	r.HandleFunc("/", landing.PageHandler).Methods("GET")
	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static",
			http.FileServer(http.Dir("./static"))))
	r.PathPrefix("/vendor").Handler(
		http.StripPrefix("/vendor",
			http.FileServer(http.Dir("./static/vendor"))))
	r.Use(loggingMiddleware)
	logrus.Infof("Listening on %s", port)
	logrus.Fatalln(http.ListenAndServe(":"+port, r))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
