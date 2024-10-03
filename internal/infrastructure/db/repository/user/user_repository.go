package repository

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/repositories"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) repositories.UserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) FindAll() ([]*entities.User, error) {
	var dbUsers []db.User

	err := r.db.Preload("User").Find(&dbUsers).Error

	if err != nil {
		return nil, err
	}

	// TODO: Add mapping logic here
	//
	users := make([]*entities.User, len(dbUsers))

	for i, dbUser := range dbUsers {
		users[i] = fromDBUser(&dbUser)
	}

	return users, nil
}
