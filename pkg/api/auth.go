package api

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("TODO_PASSWORD") == "" {
			next.ServeHTTP(w, r)
			return
		}

		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, _ := jwt.Parse(cookie.Value, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("TODO_PASSWORD")), nil
		})

		if !token.Valid {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
		next.ServeHTTP(w, r)
	})
}
