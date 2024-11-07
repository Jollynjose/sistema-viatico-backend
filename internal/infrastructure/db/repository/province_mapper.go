package repository

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"github.com/google/uuid"
)

func toDBProvince(p *entities.ProvinceValidated) *db.Province {
	return &db.Province{
		Name:       p.Name,
		Identifier: p.Identifier,
		RegionID:   p.RegionID,
		Base: db.Base{
			ID:        p.Id.String(),
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		},
	}
}

func fromDBProvince(p *db.Province) *entities.Province {
	Province := &entities.Province{
		Id:         uuid.MustParse(p.ID),
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
		Name:       p.Name,
		Identifier: p.Identifier,
		RegionID:   p.RegionID,
	}

	return Province
}
