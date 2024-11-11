package repository

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"github.com/google/uuid"
)

func fromDBFuel(f *db.Fuel) *entities.Fuel {
	var history []*entities.FuelHistoryValidated

	for _, fh := range f.History {
		history = append(history, fromDBHistory(&fh))
	}

	fuel := &entities.Fuel{
		ID:        uuid.MustParse(f.ID),
		Name:      f.Name,
		Type:      f.Type,
		History:   history,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
	}

	return fuel
}

func toDBFuel(f *entities.FuelValidated) *db.Fuel {
	var history []db.FuelHistory

	for _, fh := range f.Fuel.History {
		history = append(history, *toDBHistory(fh))
	}

	fuel := &db.Fuel{
		Name:    f.Fuel.Name,
		History: history,
		Type:    f.Fuel.Type,
		Base: db.Base{
			ID:        f.Fuel.ID.String(),
			CreatedAt: f.Fuel.CreatedAt,
			UpdatedAt: f.Fuel.UpdatedAt,
		},
	}

	return fuel
}
