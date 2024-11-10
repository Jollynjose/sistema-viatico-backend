package command

import "github.com/google/uuid"

type CreateFuelHistoryCommand struct {
	Price  float64   `json:"price"`
	FuelID uuid.UUID `json:"fuel_id"`
}
