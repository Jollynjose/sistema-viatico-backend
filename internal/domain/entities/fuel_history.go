package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type FuelHistory struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Price     float64   `json:"price"`
	FuelID    uuid.UUID `json:"fuel_id"`
}

func (f *FuelHistory) validate() error {
	if f.Price == 0 {
		return errors.New("price is required")
	}

	if err := uuid.Validate(f.FuelID.String()); err != nil {
		return err
	}

	return nil
}

func NewFuelHistory(price float64, FuelID uuid.UUID) *FuelHistory {
	return &FuelHistory{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Price:     price,
		FuelID:    FuelID,
	}
}
