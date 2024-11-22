package request

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/google/uuid"
)

type CreateUserTravelHistoryRequest struct {
	UserID                 uuid.UUID `json:"user_id" gorm:"not null"`
	JobPositionHistoryID   uuid.UUID `json:"job_position_history_id" gorm:"not null"`
	TotalPrice             float64   `json:"total_price" gorm:"not null"`
	PlusPercentage         float64   `json:"plus_percentage" gorm:"not null"`
	IsLunchApplied         bool      `json:"is_lunch_applied" gorm:"not null"`
	IsBreakfastApplied     bool      `json:"is_breakfast_applied" gorm:"not null"`
	IsDinnerApplied        bool      `json:"is_dinner_applied" gorm:"not null"`
	IsAccommodationApplied bool      `json:"is_accommodation_applied" gorm:"not null"`
	PassagePrice           float64   `json:"passage_price" gorm:"not null"`
}

func (r *CreateUserTravelHistoryRequest) ToCreateUserTravelHistoryCommand() *command.CreateUserTravelHistoryCommand {
	return &command.CreateUserTravelHistoryCommand{
		UserID:                 r.UserID,
		JobPositionHistoryID:   r.JobPositionHistoryID,
		TotalPrice:             r.TotalPrice,
		PlusPercentage:         r.PlusPercentage,
		IsLunchApplied:         r.IsLunchApplied,
		IsBreakfastApplied:     r.IsBreakfastApplied,
		IsDinnerApplied:        r.IsDinnerApplied,
		IsAccommodationApplied: r.IsAccommodationApplied,
		PassagePrice:           r.PassagePrice,
	}
}
