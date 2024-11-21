package mapper

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
)

func NewTollResultFromEntity(t *entities.Toll) *common.TollResult {
	if t == nil {
		return nil
	}

	return &common.TollResult{
		ID:        t.ID,
		Price:     t.Price,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}
