package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nstoker/gorocktrack/internal/web/authentication"
	landing "github.com/nstoker/gorocktrack/internal/web/landing"
	"github.com/sirupsen/logrus"
)

// initialiseRoutes set up the routes
func (s *Server) initialiseRoutes() {

	s.staticRoutes()
	authentication.Routes(s.Router)
	s.Router.Use(loggingMiddleware)

	s.walkRoutes()
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func (s *Server) staticRoutes() {
	s.Router.HandleFunc("/", landing.PageHandler).Methods("GET")
	s.Router.PathPrefix("/static/").Handler(
		http.StripPrefix("/static",
			http.FileServer(http.Dir("./static"))))

	s.Router.PathPrefix("/vendor").Handler(
		http.StripPrefix("/vendor",
			http.FileServer(http.Dir("./static/vendor"))))
}

func (s *Server) walkRoutes() {
	s.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		template, _ := route.GetPathTemplate()
		method, _ := route.GetMethods()

		fmt.Println(template, method)
		return nil
	})
}
