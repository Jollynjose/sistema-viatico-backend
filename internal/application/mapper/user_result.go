package mapper

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
)

func NewUserResultFromValidatedEntity(user *entities.User) *common.UserResult {
	return NewUserResultFromEntity(user)
}

func NewUserResultFromEntity(user *entities.User) *common.UserResult {
	if user == nil {
		return nil
	}

	return &common.UserResult{
		ID:                      user.Id,
		CreatedAt:               user.CreatedAt,
		UpdatedAt:               user.UpdatedAt,
		FirstName:               user.FirstName,
		LastName:                user.LastName,
		Email:                   user.Email,
		Role:                    string(user.Role),
		JobPositionID:           user.JobPositionID,
		JobPosition:             NewJobPositionResultFromEntity(user.JobPosition),
		JobPostionSpecification: user.JobPostionSpecification,
	}
}
