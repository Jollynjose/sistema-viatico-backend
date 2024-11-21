package request

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/google/uuid"
)

type CreateTollRequest struct {
	Price           float64   `json:"price"`
	Order           int       `json:"order"`
	TravelExpenseID uuid.UUID `json:"travel_expense_id"`
}

func (r *CreateTollRequest) ToCreateTollCommand() *command.CreateTollCommand {
	return &command.CreateTollCommand{
		Price:           r.Price,
		Order:           r.Order,
		TravelExpenseID: r.TravelExpenseID,
	}
}
