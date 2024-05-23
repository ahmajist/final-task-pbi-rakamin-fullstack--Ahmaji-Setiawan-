package middlewares

import (
	"context"
	"net/http"
	"strings"
	"FINALTASK-BTPN-SYARIAH/helpers"
)

// Authentication JWT
func JWTAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized: Missing authorization header"))
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		claims, err := helpers.ValidateJWT(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized: Invalid token"))
			return
		}

		// Set user information into context
		ctx := r.Context()
		ctx = context.WithValue(ctx, "username", claims.Username)
		r = r.WithContext(ctx)

		// Continue to the next handler if token is valid
		next.ServeHTTP(w, r)
	}
}
