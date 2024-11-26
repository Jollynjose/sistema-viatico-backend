package command

import (
	"time"

	"github.com/google/uuid"
)

type CreateTravelExpenseCommand struct {
	FuelHistoryID     uuid.UUID                        `json:"fuel_history_id"`
	TotalPrice        float64                          `json:"total_price"`
	DepartureDate     time.Time                        `json:"departure_date"`
	ArrivalDate       time.Time                        `json:"arrival_date"`
	SolicitudeDate    time.Time                        `json:"solicitude_date"`
	Route             CreateRouteCommand               `json:"route"`
	UserTravelHistory []CreateUserTravelHistoryCommand `json:"user_travel_history"`
	TransportType     string                           `json:"transport_type"`
	VisitMotivation   string                           `json:"visit_motivation"`
	Toll              []CreateTollCommand              `json:"toll"`
	Dependency        string                           `json:"dependency"`
}
