package request

import (
	"errors"
	"time"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/helpers"
	"github.com/google/uuid"
)

type CreateTravelExpenseRequest struct {
	ID                uuid.UUID                        `json:"id"`
	FuelHistoryID     uuid.UUID                        `json:"fuel_history_id"`
	TotalPrice        float64                          `json:"total_price"`
	CreatedAt         time.Time                        `json:"created_at"`
	UpdatedAt         time.Time                        `json:"updated_at"`
	DepartureDate     time.Time                        `json:"departure_date"`
	ArrivalDate       time.Time                        `json:"arrival_date"`
	SolicitudeDate    time.Time                        `json:"solicitude_date"`
	Route             CreateRouteRequest               `json:"route"`
	TransportType     string                           `json:"transport_type"`
	VisitMotivation   string                           `json:"visit_motivation"`
	UserTravelHistory []CreateUserTravelHistoryRequest `json:"user_travel_history"`
	Toll              []CreateTollRequest              `json:"toll"`
}

func (r *CreateTravelExpenseRequest) ToCreateTravelExpenseCommand() (*command.CreateTravelExpenseCommand, error) {
	route := *r.Route.ToCreateRouteCommand()

	travelExpense := &command.CreateTravelExpenseCommand{
		ID:                r.ID,
		FuelHistoryID:     r.FuelHistoryID,
		TotalPrice:        r.TotalPrice,
		CreatedAt:         r.CreatedAt,
		UpdatedAt:         r.UpdatedAt,
		DepartureDate:     r.DepartureDate,
		ArrivalDate:       r.ArrivalDate,
		SolicitudeDate:    r.SolicitudeDate,
		Route:             route,
		TransportType:     r.TransportType,
		VisitMotivation:   r.VisitMotivation,
		UserTravelHistory: make([]command.CreateUserTravelHistoryCommand, 0),
		Toll:              make([]command.CreateTollCommand, 0),
	}

	if helpers.IsArrayEmpty(r.UserTravelHistory) {
		return nil, errors.New("user travel history must not be empty")
	}

	if helpers.IsArrayEmpty(r.Toll) {
		return nil, errors.New("toll must not be empty")
	}

	for _, userTravelHistory := range r.UserTravelHistory {
		travelExpense.UserTravelHistory = append(travelExpense.UserTravelHistory, *userTravelHistory.ToCreateUserTravelHistoryCommand())
	}

	for _, toll := range r.Toll {
		travelExpense.Toll = append(travelExpense.Toll, *toll.ToCreateTollCommand())
	}

	return travelExpense, nil
}
