package repositories

import "github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"

type ProvinceRepository interface {
	Create(province *entities.ProvinceValidated) (*entities.Province, error)
	FindById(id string) (*entities.Province, error)
	FindByIdentifier(identifier string) (*entities.Province, error)
	FindAll() ([]*entities.Province, error)
}
