package interfaces

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
)

type UserService interface {
	Signup(userCommand *command.CreateUserCommand) (*command.CreateUserCommandResult, error)
	SignIn(userCommand *command.FindUserCommand) (*command.CreateUserCommandResult, error)
}
