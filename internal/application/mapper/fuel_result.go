package mapper

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
)

func NewFuelResultFromValidatedEntity(j *entities.FuelValidated) *common.FuelResult {
	return NewFuelResultFromEntity(&j.Fuel)
}

func NewFuelResultFromEntity(j *entities.Fuel) *common.FuelResult {
	if j == nil {
		return nil
	}

	var fuelHistoryResults []common.FuelHistoryResult

	for _, jph := range j.History {
		fuelHistoryResults = append(fuelHistoryResults, *NewFuelHistoryResultFromEntity(&jph.FuelHistory))
	}

	return &common.FuelResult{
		ID:        j.ID,
		Name:      j.Name,
		Type:      j.Type,
		CreatedAt: j.CreatedAt,
		UpdatedAt: j.UpdatedAt,
		History:   fuelHistoryResults,
	}
}
