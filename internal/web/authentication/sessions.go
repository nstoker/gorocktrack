package authentication

import (
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/nstoker/gorocktrack/internal/pkg/localenv"
)

var (
	cookieHandler = securecookie.New(
		[]byte(localenv.Hash), []byte(localenv.Block))
)

func clearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)
}

func setSession(email string, w http.ResponseWriter) {
	value := map[string]string{
		"name": email,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}

		http.SetCookie(w, cookie)
	}
}
