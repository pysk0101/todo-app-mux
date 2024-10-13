package api

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/pysk0101/todo-app-mux/backend/internal/core/ports"
)

// JWTMiddleware JWT doğrulama middleware'i
func JWTMiddleware(authService ports.AuthService) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Authorization başlığından token alın
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header missing", http.StatusUnauthorized)
				return
			}

			// Tokeni ayrıştır
			tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
			claims := &jwt.StandardClaims{}

			secret := os.Getenv("JWT_SECRET")

			// Tokeni doğrula
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil // JWT secret
			})

			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// İstek konteksine kullanıcı ID'sini ekle
			r.Header.Set("UserID", claims.Subject)

			// İsteği bir sonraki middleware'e veya handler'a yönlendir
			next.ServeHTTP(w, r)
		})
	}
}
