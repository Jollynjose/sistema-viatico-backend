package services

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/mapper"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/query"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/repositories"
	"github.com/google/uuid"
)

type FuelHistoryService struct {
	repo repositories.FuelHistoryRepository
}

func NewFuelHistoryServicee(repo repositories.FuelHistoryRepository) interfaces.FuelHistoryService {
	return &FuelHistoryService{
		repo: repo,
	}
}

func (s *FuelHistoryService) FindAllByFuelID(fuelID uuid.UUID) (*query.FuelHistoriesQueryResult, error) {
	fuels, err := s.repo.FindAllByFuelID(fuelID)
	if err != nil {
		return nil, err
	}

	var mappedFuelsHistory []*common.FuelHistoryResult

	for _, fuel := range fuels {
		mappedFuelHistory := mapper.NewFuelHistoryResultFromEntity(fuel)
		mappedFuelsHistory = append(mappedFuelsHistory, mappedFuelHistory)
	}

	return &query.FuelHistoriesQueryResult{
		Results: mappedFuelsHistory,
	}, nil
}

func (s *FuelHistoryService) FindByFuelID(FueldID uuid.UUID) (*query.FuelHistoryQueryResult, error) {
	fuelHistory, err := s.repo.FindCurrenthistory(FueldID)
	if err != nil {
		return nil, err
	}

	mappedHistory := mapper.NewFuelHistoryResultFromEntity(fuelHistory)

	return &query.FuelHistoryQueryResult{
		Result: mappedHistory,
	}, nil
}

func (s *FuelHistoryService) Create(history *command.CreateFuelHistoryCommand) (*command.CreateFuelHistoryCommandResults, error) {
	fuelHistoryEntity := entities.NewFuelHistory(
		history.Price,
		history.FuelID,
	)

	fuelHistoryEntityValidated := entities.NewFuelHistoryValidated(fuelHistoryEntity)

	if fuelHistoryEntityValidated.Err != nil {
		return nil, fuelHistoryEntityValidated.Err
	}

	result, err := s.repo.Create(fuelHistoryEntityValidated)

	if err != nil {
		return nil, err
	}

	mappedResult := mapper.NewFuelHistoryResultFromEntity(result)

	return &command.CreateFuelHistoryCommandResults{
		Result: mappedResult,
	}, nil
}
