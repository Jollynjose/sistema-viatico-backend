package repositories

import "github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"

type RegionRepository interface {
	Create(region *entities.RegionValidated) (*entities.Region, error)
	FindById(id string) (*entities.Region, error)
	FindByIdentifier(identifier string) (*entities.Region, error)
	FindAll() ([]*entities.Region, error)
}
