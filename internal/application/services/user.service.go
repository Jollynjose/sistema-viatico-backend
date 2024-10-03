package services

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/repositories"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(
	userRepository repositories.UserRepository,
) interfaces.UserService {

	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) Signup(userCommand *command.CreateUserCommand) (*common.UserResult, error) {
	return nil, nil
}
