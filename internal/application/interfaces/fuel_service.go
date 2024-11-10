package interfaces

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/query"
	"github.com/google/uuid"
)

type FuelService interface {
	FindAll() (*query.FuelsQueryResult, error)
	FindByID(ID uuid.UUID) (*query.FuelQueryResult, error)
	Create(fuel *command.CreateFuelCommand) (*command.CreateFuelCommandResult, error)
}
