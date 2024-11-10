package common

import (
	"time"

	"github.com/google/uuid"
)

type JobPositionHistoryResult struct {
	ID            uuid.UUID `json:"id"`
	Lunch         float64   `json:"lunch"`
	BreakFast     float64   `json:"breakfast"`
	Dinner        float64   `json:"dinner"`
	JobPositionID uuid.UUID `json:"job_position_id"`
	Accommodation float64   `json:"accommodation"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type JobPositionResult struct {
	ID                   uuid.UUID                  `json:"id"`
	Name                 string                     `json:"name"`
	CreatedAt            time.Time                  `json:"created_at"`
	UpdatedAt            time.Time                  `json:"updated_at"`
	JobPositionHistories []JobPositionHistoryResult `json:"job_position_histories"`
}
