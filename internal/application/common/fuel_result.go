package common

import (
	"time"

	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"github.com/google/uuid"
)

type FuelResult struct {
	ID        uuid.UUID           `json:"id"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
	Type      db.FuelPriceType    `json:"type"`
	Name      string              `json:"name"`
	History   []FuelHistoryResult `json:"fuel_histories"`
}
