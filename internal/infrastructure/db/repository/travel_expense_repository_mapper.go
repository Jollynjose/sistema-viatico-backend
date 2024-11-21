package repository

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"github.com/google/uuid"
)

func ToDBUserTravel(userTravel *entities.TravelExpense) *db.TravelExpense {
	var userTravelHistoryDB []db.UserTravelHistory
	var tolls []db.Toll

	for _, userTravelHistory := range userTravel.UserTravelHistory {
		u := ToDBUserTravelHistory(&userTravelHistory)
		userTravelHistoryDB = append(userTravelHistoryDB, *u)
	}

	for _, t := range userTravel.Toll {
		tolls = append(tolls, *ToDBToll(&t))
	}

	return &db.TravelExpense{
		Base: db.Base{
			ID:        userTravel.ID.String(),
			CreatedAt: userTravel.CreatedAt,
			UpdatedAt: userTravel.UpdatedAt,
		},
		FuelHistoryID:     userTravel.FuelHistoryID.String(),
		RouteID:           userTravel.Route.ID.String(),
		DepartureDate:     userTravel.DepartureDate,
		ArrivalDate:       userTravel.ArrivalDate,
		SolicitudeDate:    userTravel.SolicitudeDate,
		UserTravelHistory: userTravelHistoryDB,
		TransporteType:    userTravel.TransportType,
		VisitMotivation:   userTravel.VisitMotivation,
		Toll:              tolls,
		TotalPrice:        userTravel.TotalPrice,
	}
}

func FromDBUserTravel(userTravel *db.TravelExpense) *entities.TravelExpense {
	var tolls []entities.Toll
	var userTravelHistories []entities.UserTravelHistory

	for _, t := range userTravel.Toll {
		tolls = append(tolls, *FromDBToll(&t))
	}

	for _, userTravelHistory := range userTravel.UserTravelHistory {
		u := FromDBUserTravelHistory(&userTravelHistory)
		userTravelHistories = append(userTravelHistories, *u)
	}

	return &entities.TravelExpense{
		ID:                uuid.MustParse(userTravel.ID),
		FuelHistoryID:     uuid.MustParse(userTravel.FuelHistoryID),
		Route:             *FromDBRoute(&userTravel.Route),
		DepartureDate:     userTravel.DepartureDate,
		ArrivalDate:       userTravel.ArrivalDate,
		SolicitudeDate:    userTravel.SolicitudeDate,
		CreatedAt:         userTravel.CreatedAt,
		UpdatedAt:         userTravel.UpdatedAt,
		TransportType:     userTravel.TransporteType,
		VisitMotivation:   userTravel.VisitMotivation,
		Toll:              tolls,
		TotalPrice:        userTravel.TotalPrice,
		UserTravelHistory: userTravelHistories,
	}
}
