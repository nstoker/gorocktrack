package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"github.com/nstoker/gorocktrack/app"
	"github.com/nstoker/gorocktrack/config"
	"github.com/nstoker/gorocktrack/model"
	"github.com/sirupsen/logrus"
)

func main() {
	config.SetEnvironmentFile(".env")
	config.SetLogger()

	logrus.Infof("Version %s starting", app.Version)
	config.InitDatabase()
	waitTime := time.Second * 15
	app.TPL = template.Must(template.ParseGlob("templates/*.gohtml"))

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	r.HandleFunc("/", index)
	r.HandleFunc("/list", model.UserIndex)

	addr := "0.0.0.0"
	port := "3000"
	srv := &http.Server{
		Handler:      r,
		Addr:         addr + ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logrus.Error(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), waitTime)
	defer cancel()
	srv.Shutdown(ctx)
	logrus.Println("shutting down")
	os.Exit(0)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/list", http.StatusSeeOther)
}
