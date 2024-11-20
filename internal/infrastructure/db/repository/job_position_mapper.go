package repository

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"github.com/google/uuid"
)

func fromDbJobPosition(j *db.JobPosition) *entities.JobPosition {
	var jobPositionHistories []entities.JobPositionHistory

	if len(j.JobPositionHistories) > 0 {
		for _, jh := range j.JobPositionHistories {
			jobPositionHistories = append(jobPositionHistories, *fromDBJobPositionHistory(&jh))
		}
	}

	jobPosition := &entities.JobPosition{
		ID:                   uuid.MustParse(j.ID),
		CreatedAt:            j.CreatedAt,
		UpdatedAt:            j.UpdatedAt,
		Name:                 j.Name,
		JobPositionHistories: jobPositionHistories,
	}

	return jobPosition
}

func toDbJobPosition(j *entities.JobPositionValidated) *db.JobPosition {
	var jobPositionHistories []db.JobPositionHistory

	for _, jh := range j.JobPosition.JobPositionHistories {
		jobPositionHistories = append(jobPositionHistories, *toDBJobPositionHistory(&jh))
	}

	jobPosition := &db.JobPosition{
		Name:                 j.JobPosition.Name,
		JobPositionHistories: jobPositionHistories,
		Base: db.Base{
			ID:        j.JobPosition.ID.String(),
			CreatedAt: j.JobPosition.CreatedAt,
			UpdatedAt: j.JobPosition.UpdatedAt,
		},
	}

	return jobPosition
}
