package command

import "github.com/google/uuid"

type CreateTollCommand struct {
	Price           float64   `json:"price"`
	Order           int       `json:"order"`
	TravelExpenseID uuid.UUID `json:"travel_expense_id"`
}
