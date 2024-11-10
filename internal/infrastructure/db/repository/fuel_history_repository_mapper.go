package repository

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"github.com/google/uuid"
)

func fromDBHistory(f *db.FuelHistory) *entities.FuelHistoryValidated {
	history := &entities.FuelHistory{
		ID:        uuid.MustParse(f.ID),
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
		FuelID:    uuid.MustParse(f.FuelID),
		Price:     f.Price,
	}

	return entities.NewFuelHistoryValidated(history)
}

func toDBHistory(f *entities.FuelHistoryValidated) *db.FuelHistory {
	history := &db.FuelHistory{
		FuelID: f.FuelHistory.FuelID.String(),
		Price:  f.FuelHistory.Price,
		Base: db.Base{
			ID:        f.FuelHistory.ID.String(),
			CreatedAt: f.FuelHistory.CreatedAt,
			UpdatedAt: f.FuelHistory.UpdatedAt,
		},
	}

	return history
}
