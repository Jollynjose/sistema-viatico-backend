package entities

import (
	"errors"
	"time"

	"github.com/Jollynjose/sistema-viatico-backend/internal/helpers"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"github.com/google/uuid"
)

type Fuel struct {
	ID        uuid.UUID               `json:"id"`
	CreatedAt time.Time               `json:"created_at"`
	UpdatedAt time.Time               `json:"updated_at"`
	Type      db.FuelPriceType        `json:"type"`
	Name      string                  `json:"name"`
	History   []*FuelHistoryValidated `json:"fuel_histories"`
}

func (f *Fuel) AddHistory(fh *FuelHistoryValidated) {
	f.History = append(f.History, fh)
}

func (f *Fuel) validate() error {
	if helpers.IsEmpty(f.Name) {
		return errors.New("name is required")
	}

	if !helpers.IsValidFuelType(string(f.Type)) {
		return errors.New("invalid fuel type")
	}

	return nil
}

func NewFuel(name string, fuelType db.FuelPriceType) *Fuel {
	return &Fuel{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Type:      fuelType,
		Name:      name,
	}
}
