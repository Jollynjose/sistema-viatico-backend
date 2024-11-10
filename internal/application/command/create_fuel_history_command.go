package command

import "github.com/google/uuid"

type CreatehistoryCommand struct {
	Price  float64   `json:"price"`
	FuelID uuid.UUID `json:"fuel_id"`
}
