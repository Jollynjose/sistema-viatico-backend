package entities

import (
	"errors"
	"strings"
	"time"

	"github.com/Jollynjose/sistema-viatico-backend/internal/helpers"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"github.com/google/uuid"
)

type User struct {
	Id                      uuid.UUID
	CreatedAt               time.Time
	UpdatedAt               time.Time
	FirstName               string
	LastName                string
	Email                   string
	Password                string
	JobPositionID           string
	Role                    db.Roles
	JobPostionSpecification *string `json:"job_position_specification"`
}

func (u *User) validate() error {
	if helpers.IsEmpty(u.FirstName) {
		return errors.New("first name is required")
	}
	if helpers.IsEmpty(u.LastName) {
		return errors.New("last name is required")
	}
	if helpers.IsEmpty(u.Email) {
		return errors.New("email is required")
	}

	if helpers.IsEmpty(u.Email) {
		return errors.New("password is required")
	}

	if err := uuid.Validate(u.JobPositionID); err != nil {
		return errors.New("job position id is not valid")
	}

	if helpers.IsEmpty(u.Role.String()) {
		return errors.New("role is required")
	}

	if !helpers.IsValidRole(u.Role.String()) {
		return errors.New("role is not valid")
	}

	return nil
}

func NewUser(
	firstName,
	lastName,
	email,
	password,
	JobPositionId string,
	Role db.Roles,
	JobPostionSpecification *string,
) *User {
	user := &User{
		Id:                      uuid.New(),
		CreatedAt:               time.Now(),
		UpdatedAt:               time.Now(),
		FirstName:               firstName,
		LastName:                lastName,
		Email:                   strings.ToLower(email),
		Password:                password,
		Role:                    Role,
		JobPositionID:           JobPositionId,
		JobPostionSpecification: JobPostionSpecification,
	}
	return user
}
