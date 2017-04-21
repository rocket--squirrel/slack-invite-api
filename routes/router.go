package routes

import (
	"encoding/base64"
	"net/http"
	"strings"
	"time"

	"github.com/didip/tollbooth"
	"github.com/gorilla/mux"
	"github.com/trickierstinky/slack-invite-api/data"
	"github.com/trickierstinky/slack-invite-api/logs"
)

func Router() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		var handlerFunc http.HandlerFunc

		if route.Authenitcation {
			handlerFunc = BasicAuth(route.HandlerFunc)
		} else {
			handlerFunc = route.HandlerFunc
		}

		handlerFunc = SecureJSONHeaders(handlerFunc)

		handler = tollbooth.LimitFuncHandler(tollbooth.NewLimiter(1, time.Second), handlerFunc)

		handler = logs.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}

func SecureJSONHeaders(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-XSS-Protection", "1;mode=blockFilter")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("X-Content-Type-Options", "nosniff")
		w.Header().Add("X-Frame-Options", "DENY")

		next(w, r)
	}
}

func BasicAuth(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !data.ValidateUser(pair[0], pair[1]) {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}

}
