package repository

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"github.com/google/uuid"
)

func ToDBRoute(r *entities.RouteValidated) *db.Route {
	return &db.Route{
		Description:                r.Description,
		StartingPointProvinceID:    r.StartingPointProvinceID.String(),
		FinalDestinationProvinceID: r.FinalDestinationProvinceID.String(),
		Base: db.Base{
			ID:        r.ID.String(),
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		},
	}
}

func FromDBRoute(r *db.Route) *entities.Route {
	route := &entities.Route{
		ID:                         uuid.MustParse(r.ID),
		CreatedAt:                  r.CreatedAt,
		UpdatedAt:                  r.UpdatedAt,
		Description:                r.Description,
		StartingPointProvinceID:    uuid.MustParse(r.StartingPointProvinceID),
		FinalDestinationProvinceID: uuid.MustParse(r.FinalDestinationProvinceID),
	}

	return route
}
