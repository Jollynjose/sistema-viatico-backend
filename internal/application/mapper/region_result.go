package mapper

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
)

func NewRegionResultFromValidatedEntity(r *entities.Region) *common.RegionResult {
	return NewRegionResultFromEntity(r)
}

func NewRegionResultFromEntity(r *entities.Region) *common.RegionResult {
	if r == nil {
		return nil
	}

	return &common.RegionResult{
		ID:         r.Id,
		Name:       r.Name,
		Identifier: r.Identifier,
		CreatedAt:  r.CreatedAt,
		UpdatedAt:  r.UpdatedAt,
	}
}
