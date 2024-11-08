package repositories

import "github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"

type UserRepository interface {
	FindAll() ([]*entities.User, error)
	IsExist(email string) bool
	Create(user *entities.UserValidated) (*entities.User, error)
	FindOneByEmail(email string) (*entities.User, error)
	FindOneById(id string) (*entities.User, error)
}
