package services

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/mapper"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/query"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/repositories"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
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
	fuels, err := s.repo.FindAll()

	if err != nil {
		return nil, err
	}

	var mappedFuels []*common.FuelResult

	for _, fuel := range fuels {
		mappedFuel := mapper.NewFuelResultFromEntity(fuel)
		mappedFuels = append(mappedFuels, mappedFuel)
	}

	return &query.FuelsQueryResult{
		Results: mappedFuels,
	}, nil
}

func (s *FuelService) FindByID(ID uuid.UUID) (*query.FuelQueryResult, error) {
	fuel, err := s.repo.FindByID(ID)

	if err != nil {
		return nil, err
	}

	mappedFuel := mapper.NewFuelResultFromEntity(fuel)

	return &query.FuelQueryResult{
		Result: mappedFuel,
	}, nil
}

func (s *FuelService) Create(fuel *command.CreateFuelCommand) (*command.CreateFuelCommandResult, error) {
	fuelEntity := entities.NewFuel(
		fuel.Name,
		fuel.Type,
	)

	for _, history := range fuel.History {
		fuelHistory := entities.NewFuelHistory(
			history.Price,
			fuelEntity.ID,
		)

		fuelHistoryValidated := entities.NewFuelHistoryValidated(fuelHistory)

		if fuelHistoryValidated.Err != nil {
			return nil, fuelHistoryValidated.Err
		}
		fuelEntity.AddHistory(fuelHistoryValidated)
	}

	fuelEntityValidate := entities.NewFuelValidated(fuelEntity)

	if fuelEntityValidate.Err != nil {
		return nil, fuelEntityValidate.Err
	}

	result, err := s.repo.Create(fuelEntityValidate)

	if err != nil {
		return nil, err
	}

	fuelMapped := mapper.NewFuelResultFromEntity(result)

	return &command.CreateFuelCommandResult{
		Result: fuelMapped,
	}, nil
}

func (s *FuelService) FindOneByFuelHistoryId(fuelHistoryId uuid.UUID) (*db.Fuel, error) {
	fuel, err := s.repo.FindOneByFuelHistoryId(fuelHistoryId)

	if err != nil {
		return nil, err
	}

	return fuel, nil
}
