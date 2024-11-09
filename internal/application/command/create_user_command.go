package command

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"github.com/google/uuid"
)

type CreateUserCommand struct {
	ID                      uuid.UUID
	FirstName               string
	LastName                string
	Email                   string
	Password                string
	Role                    db.Roles
	JobPositionID           string
	JobPostionSpecification *string
}
