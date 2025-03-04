package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"github.com/mnshah219/go_net_http/auth/utils"
)

type UserIDKey string

func AuthMiddleware(next http.Handler) http.Handler {
	slog.Info("Intitializing auth middleware")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}
		jwt, found := strings.CutPrefix(token, "Bearer ")
		if !found {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		userID, err := utils.VerifyJWT(jwt)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserIDKey("userID"), userID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
