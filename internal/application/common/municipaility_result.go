package common

import (
	"time"

	"github.com/google/uuid"
)

type MunicipalityResult struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string    `json:"name"`
	Identifier   string    `json:"identifier"`
	RegionCode   string    `json:"region_code"`
	ProvinceCode string    `json:"province_code"`
	Code         string    `json:"code"`
}
