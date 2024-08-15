package middlewares

import (
	"access-point/config"
	"access-point/web/model"
	"log/slog"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := extractToken(r)
		if tokenString == "" {
			http.Error(w, "Authorization token is missing", http.StatusUnauthorized)
			return
		}

		claims := &model.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetConfig().JwtSecret), nil
		})
		if err != nil {
			slog.Error("Invalid token", "err", err)
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			slog.Error("expired token", "err", err)
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func extractToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, "Bearer ")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}
