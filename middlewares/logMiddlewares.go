package middlewares

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// Middleware to log incoming requests
func LogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		logrus.Infof("Request received: %s %s", r.Method, r.URL.Path)

		// Call the next handler
		next(w, r)
	}
}
