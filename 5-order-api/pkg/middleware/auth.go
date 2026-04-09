package middleware

import (
	"5-project/configs"
	"5-project/pkg/jwt"
	"fmt"
	"net/http"
	"strings"
)

func Auth(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHandler := r.Header.Get("Authorization")
		token := strings.Trim(authHandler, "Bearer")
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		fmt.Println(isValid)
		fmt.Println(data)
		next.ServeHTTP(w, r)
	})
}