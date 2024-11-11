package request

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
)

type CreateFuelHistoryCommand struct {
	Price float64 `json:"price"`
}

type CreateFuelRequest struct {
	Type    db.FuelPriceType           `json:"type"`
	Name    string                     `json:"name"`
	History []CreateFuelHistoryCommand `json:"fuel_histories"`
}

func (req *CreateFuelRequest) ToCreateFuelCommand() *command.CreateFuelCommand {
	histories := make([]command.CreateFuelHistoryCommand, 0)
	for _, history := range req.History {
		histories = append(histories, command.CreateFuelHistoryCommand{
			Price: history.Price,
		})
	}

	return &command.CreateFuelCommand{
		Type:    req.Type,
		Name:    req.Name,
		History: histories,
	}
}
