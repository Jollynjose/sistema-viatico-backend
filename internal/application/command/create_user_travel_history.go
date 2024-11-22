package command

import "github.com/google/uuid"

type CreateUserTravelHistoryCommand struct {
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
