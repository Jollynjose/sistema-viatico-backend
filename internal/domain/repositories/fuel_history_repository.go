package repositories

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/google/uuid"
)

type FuelHistoryRepository interface {
	FindAllByFuelID(fuelID uuid.UUID) ([]*entities.FuelHistory, error)
	Create(f *entities.FuelHistoryValidated) (*entities.FuelHistory, error)
	FindCurrenthistory(fuelID uuid.UUID) (*entities.FuelHistory, error)
}
