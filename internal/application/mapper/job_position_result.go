package mapper

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
)

func NewJobPositionResultFromValidatedEntity(j *entities.JobPosition) *common.JobPositionResult {
	return NewJobPositionResultFromEntity(j)
}

func NewJobPositionResultFromEntity(j *entities.JobPosition) *common.JobPositionResult {
	if j == nil {
		return nil
	}

	var jobPositionHistoryResults []common.JobPositionHistoryResult

	for _, jph := range j.JobPositionHistories {
		jobPositionHistoryResults = append(jobPositionHistoryResults, *NewJobPositionHistoryResultFromEntity(&jph))
	}

	return &common.JobPositionResult{
		ID:                   j.ID,
		Name:                 j.Name,
		CreatedAt:            j.CreatedAt,
		UpdatedAt:            j.UpdatedAt,
		JobPositionHistories: jobPositionHistoryResults,
	}
}
