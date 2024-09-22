package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (u *User) validate() error {
	if u.FirstName == "" {
		return errors.New("first name is required")
	}
	if u.LastName == "" {
		return errors.New("last name is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	return nil
}
