package middlewares

import (
	"context"
	"net/http"

	"github.com/eunkenlow/golang-api-boilerplate/pkg/firebase"
)

// IsAuthorized verify if firebase access token is valid
func IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		token, err := firebase.VerifyToken(r.Context(), reqToken)
		if err != nil {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		ctx := context.WithValue(r.Context(), "token", token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
