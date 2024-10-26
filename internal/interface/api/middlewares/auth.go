package middlewares

import (
	"context"
	"net/http"
	"os"

	"github.com/Jollynjose/sistema-viatico-backend/internal/helpers"
)

func CheckAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secretKey := os.Getenv("SECRET_KEY")

		if secretKey == "" {
			http.Error(w, "Invalid secret key", http.StatusUnauthorized)
			return
		}

		// Get the token from the request
		header := r.Header.Get("Authorization")

		if header == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		// Get the token from the request
		tokenString := header[len("Bearer "):]

		// Verify the token
		tokenVerified, error := helpers.VerifyToken(tokenString, secretKey)

		if error != nil {
			http.Error(w, error.Error(), http.StatusUnauthorized)
			return
		}

		// Save in the context
		sub, err := helpers.ParseToken(tokenVerified, secretKey)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Save in the context
		ctx := r.Context()
		ctx = context.WithValue(ctx, "userId", sub)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
