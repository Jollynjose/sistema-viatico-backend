package mapper

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/interface/api/dto/response"
)

func ToAuthUserResponse(u *common.UserResult, token string) *response.AuthUserResponse {
	return &response.AuthUserResponse{
		UserId: u.ID.String(),
		Token:  token,
	}
}
