package services

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/mapper"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/query"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/repositories"
	"github.com/google/uuid"
)

type TravelExpenseService struct {
	repo repositories.TravelExpenseRepository
}

func NewTravelExpenseService(repo repositories.TravelExpenseRepository) interfaces.TravelExpenseService {
	return &TravelExpenseService{repo: repo}
}

func (s *TravelExpenseService) Create(travel *command.CreateTravelExpenseCommand) (*command.CreateTravelExpenseCommandResult, error) {
	route := entities.NewRoute(
		travel.Route.StartingPointProvinceID,
		travel.Route.FinalDestinationProvinceID,
		travel.Route.Description,
		travel.Route.TotalKms,
	)

	travelExpense := entities.NewTravelExpenses(
		travel.FuelHistoryID,
		travel.TotalPrice,
		travel.DepartureDate,
		travel.ArrivalDate,
		travel.SolicitudeDate,
		*route,
		travel.TransportType,
		travel.VisitMotivation,
		travel.Dependency,
	)

	for _, userTravelHistory := range travel.UserTravelHistory {
		newUserTravelHistoryEntity := entities.NewUserTravelHistory(
			userTravelHistory.UserID,
			travelExpense.ID,
			userTravelHistory.JobPositionHistoryID,
			userTravelHistory.TotalPrice,
			userTravelHistory.PlusPercentage,
			userTravelHistory.IsLunchApplied,
			userTravelHistory.IsBreakfastApplied,
			userTravelHistory.IsDinnerApplied,
			userTravelHistory.IsAccommodationApplied,
			userTravelHistory.PassagePrice,
		)

		travelExpense.AddUserTravelHistory(*newUserTravelHistoryEntity)
	}

	for _, toll := range travel.Toll {
		newTollEntity := entities.NewToll(
			toll.Price,
			toll.Order,
			travelExpense.ID,
		)
		travelExpense.AddToll(*newTollEntity)
	}

	result, err := s.repo.Create(travelExpense)

	if err != nil {
		return nil, err
	}

	return &command.CreateTravelExpenseCommandResult{
		Result: mapper.NewTravelExpenseResultFromEntity(result),
	}, nil
}

func (s *TravelExpenseService) FindOne(id uuid.UUID) (*query.FindOneTravelExpenseQuery, error) {
	travel, err := s.repo.FindOne(id)

	if err != nil {
		return nil, err
	}

	result := mapper.NewTravelExpenseResultFromEntity(travel)

	return &query.FindOneTravelExpenseQuery{
		Result: result,
	}, nil
}
