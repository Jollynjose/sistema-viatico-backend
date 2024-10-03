package interfaces

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
)

type UserService interface {
	Signup(userCommand *command.CreateUserCommand) (*common.UserResult, error)
}
