package repositories

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"github.com/google/uuid"
)

type FuelRepository interface {
	FindAll() ([]*entities.Fuel, error)
	Create(f *entities.FuelValidated) (*entities.Fuel, error)
	FindByID(id uuid.UUID) (*entities.Fuel, error)
	FindOneByFuelHistoryId(fuelHistory uuid.UUID) (*db.Fuel, error)
}
