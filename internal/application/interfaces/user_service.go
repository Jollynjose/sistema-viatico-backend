package interfaces

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/query"
)

type UserService interface {
	Signup(userCommand *command.CreateUserCommand) (*command.CreateUserCommandResult, error)
	SignIn(userCommand *command.FindUserCommand) (*query.UserQueryResult, error)
}
