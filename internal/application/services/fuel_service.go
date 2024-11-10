package services

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/query"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/repositories"
	"github.com/google/uuid"
)

type FuelService struct {
	repo repositories.FuelRepository
}

func NewFuelService(repo repositories.FuelRepository) interfaces.FuelService {
	return &FuelService{
		repo: repo,
	}
}

func (s *FuelService) FindAll() (*query.FuelsQueryResult, error) {
	// fuels, err := s.repo.FindAll()
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (s *FuelService) FindByID(ID uuid.UUID) (*query.FuelQueryResult, error) {
	// fuel, err := s.repo.FindByID(ID)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (s *FuelService) Create(fuel *command.CreateFuelCommand) (*command.CreateFuelCommandResult, error) {
	return nil, nil
}
