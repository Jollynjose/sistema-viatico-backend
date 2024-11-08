package common

import (
	"time"

	"github.com/google/uuid"
)

type RegionResult struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Name       string    `json:"name"`
	Identifier string    `json:"identifier"`
	Code       string    `json:"code"`
}
