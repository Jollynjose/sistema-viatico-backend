package command

import "github.com/google/uuid"

type CreateUserCommand struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Password  string
}
