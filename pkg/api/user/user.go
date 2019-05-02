package user

import (
	"net/http"
	"time"

	"firebase.google.com/go/auth"
	"github.com/eunkenlow/golang-api-boilerplate/pkg/postgresql"
	"github.com/go-chi/render"
	"github.com/go-pg/pg"
)

// User structure
type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// CreateUser gets a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	db := postgresql.GetDb()

	token, ok := ctx.Value("token").(*auth.Token)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	user := &User{
		ID:    token.Subject,
		Email: token.Claims["email"].(string),
		Name:  token.Claims["name"].(string),
	}

	err := db.Insert(user)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	render.JSON(w, r, user)
}

// GetCurrentUser gets a user
func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	db := postgresql.GetDb()
	ctx := r.Context()

	token, ok := ctx.Value("token").(*auth.Token)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	user := &User{ID: token.Subject}
	err := db.Select(user)
	if err != nil {
		if err == pg.ErrNoRows {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		http.Error(w, http.StatusText(500), 500)
		return
	}

	render.JSON(w, r, user)
}
