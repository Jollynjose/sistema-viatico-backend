package mapper

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
)

func NewProvinceResultFromValidatedEntity(p *entities.Province) *common.ProvinceResult {
	return NewProvinceResultFromEntity(p)
}

func NewProvinceResultFromEntity(p *entities.Province) *common.ProvinceResult {
	if p == nil {
		return nil
	}

	return &common.ProvinceResult{
		ID:         p.Id,
		Name:       p.Name,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
		Identifier: p.Identifier,
		RegionCode: p.RegionCode,
		Code:       p.Code,
	}
}
