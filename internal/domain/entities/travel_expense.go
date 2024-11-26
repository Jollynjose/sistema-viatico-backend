package entities

import (
	"time"

	"github.com/google/uuid"
)

type TravelExpense struct {
	ID                uuid.UUID           `json:"id"`
	FuelHistoryID     uuid.UUID           `json:"fuel_history_id"`
	TotalPrice        float64             `json:"total_price"`
	CreatedAt         time.Time           `json:"created_at"`
	UpdatedAt         time.Time           `json:"updated_at"`
	DepartureDate     time.Time           `json:"departure_date"`
	ArrivalDate       time.Time           `json:"arrival_date"`
	SolicitudeDate    time.Time           `json:"solicitude_date"`
	Route             Route               `json:"route"`
	UserTravelHistory []UserTravelHistory `json:"user_travel_history"`
	TransportType     string              `json:"transport_type"`
	VisitMotivation   string              `json:"visit_motivation"`
	Toll              []Toll              `json:"toll"`
	Dependency        string              `json:"dependency"`
}

func NewTravelExpenses(
	fuelHistoryID uuid.UUID,
	totalPrice float64,
	departureDate time.Time,
	arrivalDate time.Time,
	SolicitudeDate time.Time,
	route Route,
	transportType string,
	visitMotivation string,
	dependency string,
) *TravelExpense {
	return &TravelExpense{
		ID:              uuid.New(),
		FuelHistoryID:   fuelHistoryID,
		TotalPrice:      totalPrice,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		DepartureDate:   departureDate,
		ArrivalDate:     arrivalDate,
		Route:           route,
		TransportType:   transportType,
		VisitMotivation: visitMotivation,
		SolicitudeDate:  SolicitudeDate,
		Dependency:      dependency,
	}
}

func (t *TravelExpense) AddUserTravelHistory(userTravelHistory UserTravelHistory) {
	t.UserTravelHistory = append(t.UserTravelHistory, userTravelHistory)
}

func (t *TravelExpense) AddToll(toll Toll) {
	t.Toll = append(t.Toll, toll)
}
