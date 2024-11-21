package mapper

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
)

func NewTravelExpenseResultFromEntity(t *entities.TravelExpense) *common.TravelExpenseResult {
	if t == nil {
		return nil
	}

	var userTravelHistories []common.UserTravelHistoryResult
	var tolls []common.TollResult

	for _, u := range t.UserTravelHistory {
		userTravelHistories = append(userTravelHistories, *NewUserTravelHistoryResultFromEntity(&u))
	}

	for _, toll := range t.Toll {
		tolls = append(tolls, *NewTollResultFromEntity(&toll))
	}

	return &common.TravelExpenseResult{
		ID:                t.ID,
		FuelHistoryID:     t.FuelHistoryID,
		TotalPrice:        t.TotalPrice,
		DepartureDate:     t.DepartureDate,
		ArrivalDate:       t.ArrivalDate,
		SolicitudeDate:    t.SolicitudeDate,
		Route:             *NewRouteResultFromEntity(&t.Route),
		Toll:              tolls,
		TransportType:     t.TransportType,
		VisitMotivation:   t.VisitMotivation,
		UserTravelHistory: userTravelHistories,
		CreatedAt:         t.CreatedAt,
		UpdatedAt:         t.UpdatedAt,
	}
}
