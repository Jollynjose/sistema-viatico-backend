package repository

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"github.com/google/uuid"
)

func toDBUser(user *entities.UserValidated) *db.User {
	return &db.User{
		FirstName:               user.FirstName,
		LastName:                user.LastName,
		Email:                   user.Email,
		Password:                user.Password,
		Role:                    user.Role,
		JobPositionID:           user.JobPositionID,
		JobPostionSpecification: user.JobPostionSpecification,
		Base: db.Base{
			ID:        user.Id.String(),
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}
}

func fromDBUser(dbUser *db.User) *entities.User {

	user := &entities.User{
		Id:                      uuid.MustParse(dbUser.ID),
		CreatedAt:               dbUser.CreatedAt,
		UpdatedAt:               dbUser.UpdatedAt,
		FirstName:               dbUser.FirstName,
		LastName:                dbUser.LastName,
		Email:                   dbUser.Email,
		Password:                dbUser.Password,
		Role:                    dbUser.Role,
		JobPositionID:           dbUser.JobPositionID,
		JobPostionSpecification: dbUser.JobPostionSpecification,
	}

	if dbUser.JobPosition != nil {
		user.JobPosition = fromDbJobPosition(dbUser.JobPosition)
	}

	return user
}
