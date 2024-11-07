package repositories

import "github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"

type MunicipalityRepository interface {
	Create(municipality *entities.MunicipalityValidated) (*entities.Municipality, error)
	FindById(id string) (*entities.Municipality, error)
	FindByIdentifier(identifier string) (*entities.Municipality, error)
	FindAll() ([]*entities.Municipality, error)
}
