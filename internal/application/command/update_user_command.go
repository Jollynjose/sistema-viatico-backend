package command

import "github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"

type UpdateUserCommand struct {
	FirstName               *string
	LastName                *string
	Email                   *string
	Password                *string
	Role                    *db.Roles
	JobPositionID           *string
	JobPostionSpecification *string
}
