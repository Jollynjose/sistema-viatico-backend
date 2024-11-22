package mapper

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/interface/api/dto/response"
)

func ToFindUserResponse(u *common.UserResult) *response.FindUserResponse {

	return &response.FindUserResponse{
		ID:                      u.ID.String(),
		Email:                   u.Email,
		FirstName:               u.FirstName,
		LastName:                u.LastName,
		CreatedAt:               u.CreatedAt,
		UpdatedAt:               u.UpdatedAt,
		Role:                    u.Role,
		JobPositionID:           u.JobPositionID,
		JobPostionSpecification: *u.JobPostionSpecification,
	}
}
