package services

import (
	"errors"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/repositories"
)

type JobPositionHistoriesService struct {
	repo repositories.JobPositionHistoryRepository
}

func NewJobPositionHistoriesService(repo repositories.JobPositionHistoryRepository) interfaces.JobPositionHistoriesService {
	return &JobPositionHistoriesService{
		repo: repo,
	}
}

func (jphs *JobPositionHistoriesService) CreateJobPositionHistory(jph *command.CreateJobPositionHistoryCommandByJobPositionID) (*command.CreateJobPositionHistoryCommandResult, error) {
	JobPositionHistory := entities.NewJobPositionHistoryWithJobPositionID(jph.Lunch, jph.BreakFast, jph.Dinner, jph.Accommodation, jph.JobPositionID)

	jobPositionHistoryValidated := entities.NewJobPositionHistoryValidated(JobPositionHistory)

	if !jobPositionHistoryValidated.IsValidated() {
		return nil, errors.New("job position history is not validated")
	}

	jobPositionHistoryCreated, err := jphs.repo.Create(jobPositionHistoryValidated)
	if err != nil {
		return nil, err
	}

	return &command.CreateJobPositionHistoryCommandResult{
		Result: &common.JobPositionHistoryResult{
			ID:            jobPositionHistoryCreated.ID,
			JobPositionID: jobPositionHistoryCreated.JobPositionID,
			CreatedAt:     jobPositionHistoryCreated.CreatedAt,
			UpdatedAt:     jobPositionHistoryCreated.UpdatedAt,
			Lunch:         jobPositionHistoryCreated.Lunch,
			BreakFast:     jobPositionHistoryCreated.BreakFast,
			Dinner:        jobPositionHistoryCreated.Dinner,
			Accommodation: jobPositionHistoryCreated.Accommodation,
		},
	}, nil
}
