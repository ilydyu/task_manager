package middlware

import (
	"context"
	"net/http"
	"strings"

	"github.com/ilydyu/task_manager.git/pkg/auth"
)

const UserIDContextKey = "userID"

func (m *Middleware) AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				http.Error(w, "Authorization header required", http.StatusUnauthorized)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
				return
			}

			tokenString := parts[1]

			claims, err := auth.ValidateToken(tokenString, m.secret)

			if err != nil {
				http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "userID", claims.Subject)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
