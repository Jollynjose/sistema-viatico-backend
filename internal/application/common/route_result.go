package common

import (
	"time"

	"github.com/google/uuid"
)

type RouteResult struct {
	ID                         uuid.UUID `json:"id"`
	StartingPointProvinceID    uuid.UUID `json:"starting_point_province_id"`
	FinalDestinationProvinceID uuid.UUID `json:"final_destination_province_id"`
	Description                string    `json:"description"`
	TotalKms                   int       `json:"total_kms"`
	CreatedAt                  time.Time `json:"created_at"`
	UpdatedAt                  time.Time `json:"updated_at"`
}
