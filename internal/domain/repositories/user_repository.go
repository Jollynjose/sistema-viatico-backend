package repositories

import "github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"

type UserRepository interface {
	FindAll() ([]*entities.User, error)
}
