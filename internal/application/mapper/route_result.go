package mapper

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
)

func NewRouteResultFromValidatedEntity(r *entities.RouteValidated) *common.RouteResult {
	return NewRouteResultFromEntity(&r.Route)
}

func NewRouteResultFromEntity(r *entities.Route) *common.RouteResult {
	return &common.RouteResult{
		ID:                         r.ID,
		StartingPointProvinceID:    r.StartingPointProvinceID,
		FinalDestinationProvinceID: r.FinalDestinationProvinceID,
		Description:                r.Description,
		TotalKms:                   r.TotalKms,
		CreatedAt:                  r.CreatedAt,
		UpdatedAt:                  r.UpdatedAt,
	}
}
