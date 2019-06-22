package routes

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/eunkenlow/golang-api-boilerplate/pkg/handlers/user"
)

// MeRouter current user routes
func MeRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", user.GetCurrentUser)
	return r
}

// UserRouter user routes
func UserRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", user.CreateUser)
	return r
}
