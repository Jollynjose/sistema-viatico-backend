package interfaces

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/query"
	"github.com/google/uuid"
)

type TravelExpenseService interface {
	Create(travel *command.CreateTravelExpenseCommand) (*command.CreateTravelExpenseCommandResult, error)
	FindOne(id uuid.UUID) (*query.FindOneTravelExpenseQuery, error)
}
