package repository

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"github.com/google/uuid"
)

func fromDBJobPositionHistory(j *db.JobPositionHistory) *entities.JobPositionHistory {
	jobPositionHistory := &entities.JobPositionHistory{
		ID:            uuid.MustParse(j.ID),
		CreatedAt:     j.CreatedAt,
		UpdatedAt:     j.UpdatedAt,
		JobPositionID: uuid.MustParse(j.JobPositionID),
		Lunch:         j.Lunch,
		Dinner:        j.Dinner,
		BreakFast:     j.BreakFast,
		Accommodation: j.Accommodation,
	}

	return jobPositionHistory
}

func toDBJobPositionHistory(j *entities.JobPositionHistory) *db.JobPositionHistory {
	jobPositionHistory := &db.JobPositionHistory{
		JobPositionID: j.JobPositionID.String(),
		Lunch:         j.Lunch,
		Dinner:        j.Dinner,
		BreakFast:     j.BreakFast,
		Accommodation: j.Accommodation,
		Base: db.Base{
			ID:        j.ID.String(),
			CreatedAt: j.CreatedAt,
			UpdatedAt: j.UpdatedAt,
		},
	}

	return jobPositionHistory
}
