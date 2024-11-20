package entities

import (
	"errors"
	"time"

	"github.com/Jollynjose/sistema-viatico-backend/internal/helpers"
	"github.com/google/uuid"
)

type Route struct {
	ID                         uuid.UUID `json:"id"`
	StartingPointProvinceID    uuid.UUID `json:"starting_point_province_id"`
	FinalDestinationProvinceID uuid.UUID `json:"final_destination_province_id"`
	Description                string    `json:"description"`
	TotalKms                   int       `json:"total_kms"`
	CreatedAt                  time.Time `json:"created_at"`
	UpdatedAt                  time.Time `json:"updated_at"`
}

func (r *Route) validate() error {
	if err := uuid.Validate(r.StartingPointProvinceID.String()); err != nil {
		return err
	}

	if err := uuid.Validate(r.FinalDestinationProvinceID.String()); err != nil {
		return err
	}

	if r.TotalKms <= 0 {
		return errors.New("total kms must be greater than 0")
	}

	if helpers.IsEmpty(r.Description) {
		return errors.New("description is required")
	}

	return nil
}

func NewRoute(startingPointProvinceID, finalDestinationProvinceID uuid.UUID, description string, totalKms int) *Route {
	route := &Route{
		ID:                         uuid.New(),
		StartingPointProvinceID:    startingPointProvinceID,
		FinalDestinationProvinceID: finalDestinationProvinceID,
		Description:                description,
		TotalKms:                   totalKms,
		CreatedAt:                  time.Now(),
		UpdatedAt:                  time.Now(),
	}
	return route
}
