package repository

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"github.com/google/uuid"
)

func ToDBToll(toll *entities.Toll) *db.Toll {
	return &db.Toll{
		Order:           toll.Order,
		Price:           toll.Price,
		TravelExpenseID: toll.TravelExpenseID.String(),
		Base: db.Base{
			ID:        toll.ID.String(),
			CreatedAt: toll.CreatedAt,
			UpdatedAt: toll.UpdatedAt,
		},
	}
}

func FromDBToll(toll *db.Toll) *entities.Toll {
	return &entities.Toll{
		ID:              uuid.MustParse(toll.ID),
		Order:           toll.Order,
		Price:           toll.Price,
		TravelExpenseID: uuid.MustParse(toll.TravelExpenseID),
		CreatedAt:       toll.CreatedAt,
		UpdatedAt:       toll.UpdatedAt,
	}
}
