package mapper

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
)

func NewUserTravelHistoryResultFromEntity(u *entities.UserTravelHistory) *common.UserTravelHistoryResult {
	if u == nil {
		return nil
	}

	return &common.UserTravelHistoryResult{
		ID:                     u.ID,
		UserID:                 u.UserID,
		TravelExpenseID:        u.TravelExpenseID,
		JobPositionHistoryID:   u.JobPositionHistoryID,
		TotalPrice:             u.TotalPrice,
		PlusPercentage:         u.PlusPercentage,
		IsLunchApplied:         u.IsLunchApplied,
		IsBreakfastApplied:     u.IsBreakfastApplied,
		IsDinnerApplied:        u.IsDinnerApplied,
		IsAccommodationApplied: u.IsAccommodationApplied,
		PassagePrice:           u.PassagePrice,
		CreatedAt:              u.CreatedAt,
		UpdatedAt:              u.UpdatedAt,
	}
}
