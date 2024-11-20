package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type JobPositionHistory struct {
	ID            uuid.UUID
	Lunch         float64
	BreakFast     float64
	Dinner        float64
	Accommodation float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	JobPositionID uuid.UUID
}

func (jh *JobPositionHistory) validate() error {
	if jh.Lunch < 0 {
		return errors.New("lunch must be greater than or equal to 0")
	}

	if jh.BreakFast < 0 {
		return errors.New("breakfast must be greater than or equal to 0")
	}

	if jh.Dinner < 0 {
		return errors.New("dinner must be greater than or equal to 0")
	}

	if jh.Accommodation < 0 {
		return errors.New("accommodation must be greater than or equal to 0")
	}

	return nil
}

func NewJobPositionHistory(lunch, breakfast, dinner, accommodation float64) *JobPositionHistory {
	jobPositionHistory := &JobPositionHistory{
		ID:            uuid.New(),
		Lunch:         lunch,
		BreakFast:     breakfast,
		Dinner:        dinner,
		Accommodation: accommodation,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	return jobPositionHistory
}

func NewJobPositionHistoryWithJobPositionID(lunch, breakfast, dinner, accommodation float64, jobPositionID uuid.UUID) *JobPositionHistory {
	jobPositionHistory := &JobPositionHistory{
		ID:            uuid.New(),
		Lunch:         lunch,
		BreakFast:     breakfast,
		Dinner:        dinner,
		Accommodation: accommodation,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		JobPositionID: jobPositionID,
	}

	return jobPositionHistory
}
