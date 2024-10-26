package common

import (
	"time"

	"github.com/google/uuid"
)

type UserResult struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	FirstName string
	LastName  string
	Email     string
	Password  string
}
