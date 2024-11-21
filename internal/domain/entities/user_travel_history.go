package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserTravelHistory struct {
	ID                     uuid.UUID `json:"id"`
	UserID                 uuid.UUID `json:"user_id" gorm:"not null"`
	TravelExpenseID        uuid.UUID `json:"travel_expense_id" gorm:"not null"`
	JobPositionHistoryID   uuid.UUID `json:"job_position_history_id" gorm:"not null"`
	TotalPrice             float64   `json:"total_price" gorm:"not null"`
	PlusPercentage         float64   `json:"plus_percentage" gorm:"not null"`
	IsLunchApplied         bool      `json:"is_lunch_applied" gorm:"not null"`
	IsBreakfastApplied     bool      `json:"is_breakfast_applied" gorm:"not null"`
	IsDinnerApplied        bool      `json:"is_dinner_applied" gorm:"not null"`
	IsAccommodationApplied bool      `json:"is_accommodation_applied" gorm:"not null"`
	PassagePrice           float64   `json:"passage_price" gorm:"not null"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

func NewUserTravelHistory(userID uuid.UUID, travelExpenseID uuid.UUID, jobPositionHistoryID uuid.UUID, totalPrice float64, plusPercentage float64, isLunchApplied bool, isBreakfastApplied bool, isDinnerApplied bool, isAccommodationApplied bool, passagePrice float64) *UserTravelHistory {
	return &UserTravelHistory{
		ID:                     uuid.New(),
		UserID:                 userID,
		TravelExpenseID:        travelExpenseID,
		JobPositionHistoryID:   jobPositionHistoryID,
		TotalPrice:             totalPrice,
		PlusPercentage:         plusPercentage,
		IsLunchApplied:         isLunchApplied,
		IsBreakfastApplied:     isBreakfastApplied,
		IsDinnerApplied:        isDinnerApplied,
		IsAccommodationApplied: isAccommodationApplied,
		PassagePrice:           passagePrice,
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
	}
}
