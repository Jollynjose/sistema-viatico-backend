package entities

import (
	"time"

	"github.com/google/uuid"
)

type Toll struct {
	ID              uuid.UUID `json:"id"`
	Price           float64   `json:"price"`
	Order           int       `json:"order"`
	TravelExpenseID uuid.UUID `json:"travel_expense_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func NewToll(price float64, order int, travelExpenseID uuid.UUID) *Toll {
	return &Toll{
		ID:              uuid.New(),
		Price:           price,
		Order:           order,
		TravelExpenseID: travelExpenseID,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}
