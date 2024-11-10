package common

import (
	"time"

	"github.com/google/uuid"
)

type FuelHistoryResult struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Price     float64   `json:"price"`
	FuelID    uuid.UUID `json:"fuel_id"`
}
