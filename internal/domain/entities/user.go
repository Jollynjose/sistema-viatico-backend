package entities

import (
	"errors"
	"strings"
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

func NewUser(firstName, lastName, email, password string) *User {
	user := &User{
		Id:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     strings.ToLower(email),
		Password:  password,
	}
	return user
}
