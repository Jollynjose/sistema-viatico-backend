package services

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/query"
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

func (s *FuelHistoryService) FindAll() (*query.FuelHistoriesQueryResult, error) {
	// fuels, err := s.repo.FindAll()
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (s *FuelHistoryService) FindByFuelID(ID uuid.UUID) (*query.FuelHistoryQueryResult, error) {
	// fuel, err := s.repo.FindByID(ID)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (s *FuelHistoryService) Create(history *command.CreatehistoryCommand) (*command.CreatehistoryCommandResults, error) {
	return nil, nil
}
