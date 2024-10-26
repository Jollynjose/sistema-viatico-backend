package repository

import (
	"strings"

	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/repositories"
	"github.com/Jollynjose/sistema-viatico-backend/internal/helpers"
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

	users := make([]*entities.User, len(dbUsers))

	for i, dbUser := range dbUsers {
		users[i] = fromDBUser(&dbUser)
	}

	return users, nil
}

func (r *GormUserRepository) IsExist(email string) bool {
	var user db.User

	r.db.First(&user, "email = ?", strings.ToLower(email))

	if user.ID == "" {
		return false
	}
	return true
}

func (r *GormUserRepository) Create(user *entities.UserValidated) (*entities.User, error) {
	dbUser := toDBUser(user)

	// hash password
	hashedPassword, err := helpers.GeneratePassword(dbUser.Password)

	if err != nil {
		return nil, err
	}

	dbUser.Password = hashedPassword

	if err := r.db.Create(dbUser).Error; err != nil {
		return nil, err
	}

	return fromDBUser(dbUser), nil
}

func (r *GormUserRepository) FindById(id string) (*entities.User, error) {
	var dbUser db.User

	err := r.db.Preload("User").First(&dbUser, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return fromDBUser(&dbUser), nil
}

func (r *GormUserRepository) FindOneByEmail(email string) (*entities.User, error) {
	var dbUser db.User

	err := r.db.First(&dbUser, "email = ?", strings.ToLower(email)).Error

	if err != nil {
		return nil, err
	}

	return fromDBUser(&dbUser), nil
}
