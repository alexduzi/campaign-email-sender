package endpoints

import (
	"campaignemailsender/internal/infrastructure/credentials"
	"context"
	"net/http"

	"github.com/go-chi/render"
)

type ValidateTokenFunc func(token string, ctx context.Context) (string, error)

var ValidateToken ValidateTokenFunc = credentials.ValidateToken

func Auth(nex http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": "request does not contain an authorization header"})
			return
		}

		email, err := ValidateToken(tokenString, r.Context())

		if err != nil {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), "email", email)

		nex.ServeHTTP(w, r.WithContext(ctx))
	})
}
