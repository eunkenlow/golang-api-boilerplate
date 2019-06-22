package user

import (
	"net/http"

	c "github.com/eunkenlow/golang-api-boilerplate/constants"

	e "github.com/eunkenlow/golang-api-boilerplate/pkg/apperror"
	"github.com/eunkenlow/golang-api-boilerplate/pkg/models"
	u "github.com/eunkenlow/golang-api-boilerplate/pkg/utils"

	"github.com/go-chi/render"
	"github.com/go-pg/pg"
)

// CreateUser gets a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(u.ContextKey(c.ContextUserID)).(string)
	email := r.Context().Value(u.ContextKey(c.ContextEmail)).(string)
	name := r.Context().Value(u.ContextKey(c.ContextName)).(string)

	user := &models.User{
		ID:    userID,
		Email: email,
		Name:  name,
	}

	err := user.Create()
	if err != nil {
		render.Render(w, r, e.ErrRender(err))
		return
	}

	render.JSON(w, r, user)
}

// GetCurrentUser gets a user
func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(u.ContextKey(c.ContextUserID)).(string)

	user, err := models.GetUserByID(userID)
	if err != nil {
		if err == pg.ErrNoRows {
			render.Render(w, r, e.ErrNotFound)
			return
		}
		render.Render(w, r, e.ErrRender(err))
		return
	}

	render.JSON(w, r, user)
}
