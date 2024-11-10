package mapper

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
)

func NewFuelHistoryResultFromValidatedEntity(j *entities.FuelHistoryValidated) *common.FuelHistoryResult {
	return NewFuelHistoryResultFromEntity(&j.FuelHistory)
}

func NewFuelHistoryResultFromEntity(j *entities.FuelHistory) *common.FuelHistoryResult {
	if j == nil {
		return nil
	}

	return &common.FuelHistoryResult{
		ID:        j.ID,
		FuelID:    j.FuelID,
		Price:     j.Price,
		CreatedAt: j.CreatedAt,
		UpdatedAt: j.UpdatedAt,
	}
}
