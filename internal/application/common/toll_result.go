package common

import (
	"time"

	"github.com/google/uuid"
)

type TollResult struct {
	ID              uuid.UUID `json:"id"`
	Price           float64   `json:"price"`
	Order           int       `json:"order"`
	TravelExpenseID uuid.UUID `json:"travel_expense_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
