package repository

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"github.com/google/uuid"
)

func ToDBUserTravelHistory(userTravel *entities.UserTravelHistory) *db.UserTravelHistory {
	return &db.UserTravelHistory{
		Base: db.Base{
			ID:        userTravel.ID.String(),
			CreatedAt: userTravel.CreatedAt,
			UpdatedAt: userTravel.UpdatedAt,
		},
		TravelExpenseID:        userTravel.TravelExpenseID.String(),
		JobPositionHistoryID:   userTravel.JobPositionHistoryID.String(),
		UserID:                 userTravel.UserID.String(),
		TotalPrice:             userTravel.TotalPrice,
		PlusPercentage:         userTravel.PlusPercentage,
		IsLunchApplied:         userTravel.IsLunchApplied,
		IsBreakfastApplied:     userTravel.IsBreakfastApplied,
		IsDinnerApplied:        userTravel.IsDinnerApplied,
		IsAccommodationApplied: userTravel.IsAccommodationApplied,
		PassagePrice:           userTravel.PassagePrice,
	}
}

func FromDBUserTravelHistory(userTravel *db.UserTravelHistory) *entities.UserTravelHistory {
	return &entities.UserTravelHistory{
		ID:                     uuid.MustParse(userTravel.ID),
		TravelExpenseID:        uuid.MustParse(userTravel.TravelExpenseID),
		JobPositionHistoryID:   uuid.MustParse(userTravel.JobPositionHistoryID),
		UserID:                 uuid.MustParse(userTravel.UserID),
		TotalPrice:             userTravel.TotalPrice,
		PlusPercentage:         userTravel.PlusPercentage,
		IsLunchApplied:         userTravel.IsLunchApplied,
		IsBreakfastApplied:     userTravel.IsBreakfastApplied,
		IsDinnerApplied:        userTravel.IsDinnerApplied,
		IsAccommodationApplied: userTravel.IsAccommodationApplied,
		PassagePrice:           userTravel.PassagePrice,
		CreatedAt:              userTravel.CreatedAt,
		UpdatedAt:              userTravel.UpdatedAt,
	}
}
