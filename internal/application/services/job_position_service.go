package services

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/mapper"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/query"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/repositories"
)

type JobPositionService struct {
	repo repositories.JobPositionRepository
}

func NewJobPositionService(repo repositories.JobPositionRepository) interfaces.JobPositionService {
	return &JobPositionService{
		repo: repo,
	}
}

func (s *JobPositionService) FindAll() (*query.JobPositionsQueryResult, error) {
	jobPositions, err := s.repo.FindAll()

	if err != nil {
		return nil, err
	}

	var mappedJobPositions []*common.JobPositionResult

	for _, jobPosition := range jobPositions {
		mappedJobPosition := mapper.NewJobPositionResultFromEntity(jobPosition)
		mappedJobPositions = append(mappedJobPositions, mappedJobPosition)
	}

	return &query.JobPositionsQueryResult{
		Results: mappedJobPositions,
	}, nil
}

func (s *JobPositionService) FindById(jobPositionCommand *command.FindJobPositionByIdCommand) (*query.JobPositionQueryResult, error) {
	jobPosition, err := s.repo.FindById(jobPositionCommand.ID)

	if err != nil {
		return nil, err
	}

	mappedJobPosition := mapper.NewJobPositionResultFromEntity(jobPosition)

	return &query.JobPositionQueryResult{
		Result: mappedJobPosition,
	}, nil
}

func (s *JobPositionService) Create(jobPositionCommand *command.CreateJobPositionCommand) (*command.CreateJobPositionCommandResult, error) {

	var jobPositionHistories []entities.JobPositionHistory

	for _, jph := range jobPositionCommand.JobPositionHistories {
		jobPositionHistories = append(jobPositionHistories, *entities.NewJobPositionHistory(
			jph.BreakFast,
			jph.Lunch,
			jph.Dinner,
			jph.Accommodation,
		))
	}

	newJobPosition := entities.NewJobPosition(
		jobPositionCommand.Name,
		jobPositionHistories,
	)

	newJobPositionValidated := entities.NewJobPostionValidated(newJobPosition)

	if newJobPositionValidated.Err != nil {
		return nil, newJobPositionValidated.Err
	}

	createdJobPosition, err := s.repo.Create(newJobPositionValidated)

	if err != nil {
		return nil, err
	}

	mappedJobPosition := mapper.NewJobPositionResultFromEntity(createdJobPosition)

	res := command.CreateJobPositionCommandResult{
		Result: mappedJobPosition,
	}

	return &res, nil
}
