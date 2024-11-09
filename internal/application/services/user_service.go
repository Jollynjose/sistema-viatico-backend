package services

import (
	"errors"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/mapper"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/query"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/repositories"
	"github.com/Jollynjose/sistema-viatico-backend/internal/helpers"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"github.com/google/uuid"
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
		userCommand.JobPositionID,
		userCommand.Role,
		userCommand.JobPostionSpecification,
	)

	validatedUser := entities.NewUserValidated(newUser)

	if validatedUser.Err != nil {
		return nil, validatedUser.Err
	}

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

func (u *UserService) FindUserById(userCommand *command.FindUserByIdCommand) (*query.UserQueryResult, error) {
	user, err := u.userRepository.FindOneById(userCommand.ID)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	mappedUser := mapper.NewUserResultFromValidatedEntity(user)

	res := query.UserQueryResult{
		Result: mappedUser,
	}

	return &res, nil
}

func (u *UserService) FindAll() (*query.UsersQueryResult, error) {
	users, err := u.userRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var mappedUsers []*common.UserResult

	for _, user := range users {
		mappedUser := mapper.NewUserResultFromValidatedEntity(user)

		mappedUsers = append(mappedUsers, mappedUser)
	}

	return &query.UsersQueryResult{
		Results: mappedUsers,
	}, nil
}

func (u *UserService) UpdateById(id string, userCommand *command.UpdateUserCommand) (*command.UpdateUserCommandResult, error) {
	var dbUser entities.User

	if userCommand.Email != nil {
		email := string(*userCommand.Email)

		if u.userRepository.IsExist(email) {
			return nil, errors.New("user already exists")
		}

		if helpers.IsEmpty(email) {
			return nil, errors.New("email is required")
		}
		dbUser.Email = email
	}

	if userCommand.FirstName != nil {
		firstName := string(*userCommand.FirstName)
		if helpers.IsEmpty(firstName) {
			return nil, errors.New("first name is required")
		}
		dbUser.FirstName = firstName
	}

	if userCommand.LastName != nil {
		lastName := string(*userCommand.LastName)
		if helpers.IsEmpty(lastName) {
			return nil, errors.New("last name is required")
		}
		dbUser.LastName = lastName
	}

	if userCommand.Password != nil {
		password := string(*userCommand.Password)
		if helpers.IsEmpty(password) {
			return nil, errors.New("password is required")
		}
		dbUser.Password = password
	}

	if userCommand.Role != nil {
		if !helpers.IsValidRole(string(*userCommand.Role)) {
			return nil, errors.New("role is not valid")
		}
		dbUser.Role = db.Roles(*userCommand.Role)
	}

	if userCommand.JobPositionID != nil {
		jobPositionID := string(*userCommand.JobPositionID)
		if helpers.IsEmpty(jobPositionID) {
			return nil, errors.New("job position id is required")
		}

		if err := uuid.Validate(jobPositionID); err != nil {
			return nil, err
		}

		dbUser.JobPositionID = jobPositionID
	}

	if userCommand.JobPostionSpecification != nil {
		dbUser.JobPostionSpecification = userCommand.JobPostionSpecification
	}

	u.userRepository.UpdateById(id, &dbUser)

	return nil, nil
}
