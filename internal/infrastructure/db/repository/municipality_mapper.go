package repository

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"github.com/google/uuid"
)

func toDBMunicipality(m *entities.MunicipalityValidated) *db.Municipality {
	return &db.Municipality{
		Name:       m.Name,
		Identifier: m.Identifier,
		RegionID:   m.RegionID,
		ProvinceID: m.ProvinceID,
		Base: db.Base{
			ID:        m.Id.String(),
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		},
	}
}

func fromDBMunicipality(m *db.Municipality) *entities.Municipality {
	municipality := &entities.Municipality{
		Id:         uuid.MustParse(m.ID),
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
		Name:       m.Name,
		Identifier: m.Identifier,
		RegionID:   m.RegionID,
		ProvinceID: m.ProvinceID,
	}

	return municipality
}
