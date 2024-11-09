package request

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
)

type UpdateUserByIdRequest struct {
	FirstName               *string   `json:"first_name"`
	LastName                *string   `json:"last_name"`
	Email                   *string   `json:"email"`
	Password                *string   `json:"password"`
	Role                    *db.Roles `json:"role"`
	JobPositionID           *string   `json:"job_position_id"`
	JobPostionSpecification *string   `json:"job_position_specification"`
}

func (req *UpdateUserByIdRequest) ToUpdateUserCommand() *command.UpdateUserCommand {
	return &command.UpdateUserCommand{
		FirstName:               req.FirstName,
		LastName:                req.LastName,
		Email:                   req.Email,
		Password:                req.Password,
		Role:                    req.Role,
		JobPositionID:           req.JobPositionID,
		JobPostionSpecification: req.JobPostionSpecification,
	}
}
