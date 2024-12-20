package mapper

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/interface/api/dto/response"
)

func ToFindAllUser(u *common.UserResult, jp *response.JobPosition) *response.FindAllUser {
	var jobPostionSpecification string

	if u.JobPostionSpecification != nil {
		jobPostionSpecification = *u.JobPostionSpecification
	}

	return &response.FindAllUser{
		FindUserResponse: response.FindUserResponse{
			ID:                      u.ID.String(),
			Email:                   u.Email,
			FirstName:               u.FirstName,
			LastName:                u.LastName,
			CreatedAt:               u.CreatedAt,
			UpdatedAt:               u.UpdatedAt,
			Role:                    u.Role,
			JobPositionID:           u.JobPositionID,
			JobPostionSpecification: jobPostionSpecification,
		},
		JobPosition: *jp,
	}
}
