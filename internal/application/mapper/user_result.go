package mapper

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
)

func NewUserResultFromValidatedEntity(user *entities.UserValidated) *common.UserResult {
	return NewUserResultFromEntity(&user.User)
}

func NewUserResultFromEntity(user *entities.User) *common.UserResult {
	if user == nil {
		return nil
	}

	return &common.UserResult{
		Id:        user.Id,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}
