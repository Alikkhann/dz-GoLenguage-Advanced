package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHandler := r.Header.Get("Authorization")
		token := strings.Trim(authHandler, "Bearer")
		fmt.Printf("Token: %s", token)
		next.ServeHTTP(w, r)
	})
}