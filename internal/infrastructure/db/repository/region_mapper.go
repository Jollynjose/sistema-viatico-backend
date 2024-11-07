package repository

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"github.com/google/uuid"
)

func toDBRegion(r *entities.RegionValidated) *db.Region {
	return &db.Region{
		Name:       r.Name,
		Identifier: r.Identifier,
		Base: db.Base{
			ID:        r.Id.String(),
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		},
	}
}

func fromDBRegion(r *db.Region) *entities.Region {
	region := &entities.Region{
		Id:         uuid.MustParse(r.ID),
		CreatedAt:  r.CreatedAt,
		UpdatedAt:  r.UpdatedAt,
		Name:       r.Name,
		Identifier: r.Identifier,
	}

	return region
}
