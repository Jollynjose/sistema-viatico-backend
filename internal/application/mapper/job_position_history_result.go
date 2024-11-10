package mapper

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
)

func NewJobPositionHistoryResultFromValidatedEntity(j *entities.JobPositionHistory) *common.JobPositionHistoryResult {
	return NewJobPositionHistoryResultFromEntity(j)
}

func NewJobPositionHistoryResultFromEntity(j *entities.JobPositionHistory) *common.JobPositionHistoryResult {
	if j == nil {
		return nil
	}

	return &common.JobPositionHistoryResult{
		ID:            j.ID,
		JobPositionID: j.JobPositionID,
		Lunch:         j.Lunch,
		BreakFast:     j.BreakFast,
		Dinner:        j.Dinner,
		Accommodation: j.Accommodation,
		CreatedAt:     j.CreatedAt,
		UpdatedAt:     j.UpdatedAt,
	}
}
