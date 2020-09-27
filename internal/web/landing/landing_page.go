package landing

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/sirupsen/logrus"
)

var templates = template.Must(template.ParseGlob("internal/web/landing/templates/*"))

// PageHandler displays the landing page code
func PageHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("templates: %+v", templates)
	err := templates.ExecuteTemplate(w, "landing.html", nil)
	if err != nil {
		logrus.Infof("internal/web/landing/PageHandler: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("landing page handler %s", err)))
		return
	}
}
