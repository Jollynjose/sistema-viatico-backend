package request

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
)

type SignUpUserRequest struct {
	FirstName               string   `json:"first_name"`
	LastName                string   `json:"last_name"`
	Email                   string   `json:"email"`
	Password                string   `json:"password"`
	Role                    db.Roles `json:"role"`
	JobPositionID           string   `json:"job_position_id"`
	JobPostionSpecification *string  `json:"job_position_specification"`
}

func (req *SignUpUserRequest) ToCreateUserCommand() *command.CreateUserCommand {
	return &command.CreateUserCommand{
		FirstName:               req.FirstName,
		LastName:                req.LastName,
		Email:                   req.Email,
		Password:                req.Password,
		Role:                    req.Role,
		JobPositionID:           req.JobPositionID,
		JobPostionSpecification: req.JobPostionSpecification,
	}
}
