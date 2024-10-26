package request

import "github.com/Jollynjose/sistema-viatico-backend/internal/application/command"

type SignInUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *SignInUserRequest) ToFindUserCommand() *command.FindUserCommand {
	return &command.FindUserCommand{
		Email:    req.Email,
		Password: req.Password,
	}
}
