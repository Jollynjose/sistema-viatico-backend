package interfaces

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/query"
	"github.com/google/uuid"
)

type FuelHistoryService interface {
	FindAllByFuelID(fuelID uuid.UUID) (*query.FuelHistoriesQueryResult, error)
	FindByFuelID(FuelID uuid.UUID) (*query.FuelHistoryQueryResult, error)
	Create(history *command.CreateFuelHistoryCommand) (*command.CreateFuelHistoryCommandResults, error)
}
