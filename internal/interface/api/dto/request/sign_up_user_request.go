package request

import "github.com/Jollynjose/sistema-viatico-backend/internal/application/command"

type SignUpUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (req *SignUpUserRequest) ToCreateUserCommand() *command.CreateUserCommand {
	return &command.CreateUserCommand{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}
}
