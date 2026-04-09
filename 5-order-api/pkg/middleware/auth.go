package middleware

import (
	"5-project/configs"
	"5-project/pkg/jwt"
	"5-project/pkg/resp"
	"context"
	"fmt"
	"net/http"
	"strings"
)

func Auth(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			resp.Json(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			resp.Json(w, "Invalid authorization format", http.StatusUnauthorized)
			return
		}

		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		if !isValid {
			resp.Json(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		fmt.Println(isValid)
		fmt.Println(data)
		ctx := context.WithValue(r.Context(), "user", data)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
