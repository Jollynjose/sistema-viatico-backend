package common

import (
	"time"

	"github.com/google/uuid"
)

type ProvinceResult struct {
	ID         uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Name       string
	Identifier string
	RegionID   string
}
