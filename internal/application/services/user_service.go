package services

import (
	"errors"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/mapper"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/query"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/repositories"
	"github.com/Jollynjose/sistema-viatico-backend/internal/helpers"
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

func (u *UserService) Signup(userCommand *command.CreateUserCommand) (*command.CreateUserCommandResult, error) {
	// Check if user already exists
	if u.userRepository.IsExist(userCommand.Email) {
		return nil, errors.New("user already exists")
	}

	newUser := entities.NewUser(
		userCommand.FirstName,
		userCommand.LastName,
		userCommand.Email,
		userCommand.Password,
	)

	validatedUser := entities.NewUserValidated(newUser)

	user, err := u.userRepository.Create(validatedUser)

	if err != nil {
		return nil, err
	}

	mappedUser := mapper.NewUserResultFromValidatedEntity(user)

	res := command.CreateUserCommandResult{
		Result: mappedUser,
	}

	return &res, nil
}

func (u *UserService) SignIn(userCommand *command.FindUserCommand) (*query.UserQueryResult, error) {
	user, err := u.userRepository.FindOneByEmail(userCommand.Email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	if !helpers.VerifyPassword(userCommand.Password, user.Password) {
		return nil, errors.New("invalid password")
	}

	mappedUser := mapper.NewUserResultFromValidatedEntity(user)

	res := query.UserQueryResult{
		Result: mappedUser,
	}

	return &res, nil
}
