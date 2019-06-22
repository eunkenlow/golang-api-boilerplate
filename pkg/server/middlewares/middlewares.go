package middlewares

import (
	"context"
	"net/http"

	c "github.com/eunkenlow/golang-api-boilerplate/constants"

	e "github.com/eunkenlow/golang-api-boilerplate/pkg/apperror"
	"github.com/eunkenlow/golang-api-boilerplate/pkg/firebase"
	u "github.com/eunkenlow/golang-api-boilerplate/pkg/utils"

	"github.com/go-chi/render"
)

// IsAuthorized verify if firebase access token is valid
func IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		token, err := firebase.VerifyToken(r.Context(), reqToken)
		if err != nil {
			render.Render(w, r, e.ErrUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), u.ContextKey(c.ContextUserID), token.Subject)
		ctx = context.WithValue(ctx, u.ContextKey(c.ContextEmail), token.Claims["email"])
		ctx = context.WithValue(ctx, u.ContextKey(c.ContextName), token.Claims["name"])

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
