package mapper

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
)

func NewMunicipalityResultFromValidatedEntity(e *entities.Municipality) *common.MunicipalityResult {
	return NewMunicipalityResultFromEntity(e)
}

func NewMunicipalityResultFromEntity(e *entities.Municipality) *common.MunicipalityResult {
	if e == nil {
		return nil
	}

	return &common.MunicipalityResult{
		ID:         e.Id,
		Name:       e.Name,
		CreatedAt:  e.CreatedAt,
		UpdatedAt:  e.UpdatedAt,
		Identifier: e.Identifier,
		RegionID:   e.RegionID,
		ProvinceID: e.ProvinceID,
	}
}
