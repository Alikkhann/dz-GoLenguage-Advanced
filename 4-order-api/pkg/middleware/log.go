package middleware

import (
	"net/http"
	"github.com/sirupsen/logrus"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := logrus.New()
		logger.Formatter = &logrus.JSONFormatter{}
		
		logger.WithFields(logrus.Fields{
			"method": r.Method,
			"url": r.URL.Path,
		}).Info("Incoming request")

		next.ServeHTTP(w, r)
	})	
}