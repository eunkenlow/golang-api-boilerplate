package models

import (
	"time"

	"github.com/eunkenlow/golang-api-boilerplate/pkg/postgresql"
)

// User structure
type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Create new user in db
func (user *User) Create() error {
	db := postgresql.DB()

	err := db.Insert(user)
	return err
}

// GetUserByID from db
func GetUserByID(id string) (*User, error) {
	db := postgresql.DB()

	user := &User{ID: id}
	err := db.Select(user)
	return user, err
}
