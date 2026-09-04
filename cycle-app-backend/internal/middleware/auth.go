package middleware

import (
	"context"
	"net/http"
	"strings"

	"cycle-app-backend/internal/utils"
)

type contextKey string

const userContextKey contextKey = "user"

func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		claims, err := utils.ValidateToken(tokenString)

		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		userID, ok := claims["user_id"].(float64)

		if !ok {
			http.Error(w, "Invalid user information", http.StatusUnauthorized)
			return
		}

		email, ok := claims["email"].(string)

		if !ok {
			http.Error(w, "Invalid user information", http.StatusUnauthorized)
			return
		}

		user := map[string]interface{}{
			"user_id": int(userID),
			"email":   email,
		}

		ctx := context.WithValue(r.Context(), userContextKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UserContextKey() interface{} {
	return userContextKey
}
