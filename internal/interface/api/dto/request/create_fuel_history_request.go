package request

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/google/uuid"
)

type CreateFuelHistoryRequest struct {
	Price  float64   `json:"price"`
	FuelID uuid.UUID `json:"fuel_id"`
}

func (req *CreateFuelHistoryRequest) ToCreateFuelHistoryCommand() *command.CreateFuelHistoryCommand {
	return &command.CreateFuelHistoryCommand{
		Price:  req.Price,
		FuelID: req.FuelID,
	}
}
