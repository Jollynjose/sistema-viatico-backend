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

func NewGormUserRepository(db *gorm.DB) *repositories.UserRepository {
	return &GormUserRepository{db}
}

func (r *GormUserRepository) FindAll() (*[]entities.User, error) {
	var users []db.User

	err := r.db.Preload("User").Find(&users).Error

	if err != nil {
		return nil, err
	}

	// TODO: Add mapping logic here

	return users, nil
}
