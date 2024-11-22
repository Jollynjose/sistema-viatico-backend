package request

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
)

type CreateTollRequest struct {
	Price float64 `json:"price"`
	Order int     `json:"order"`
}

func (r *CreateTollRequest) ToCreateTollCommand() *command.CreateTollCommand {
	return &command.CreateTollCommand{
		Price: r.Price,
		Order: r.Order,
	}
}
