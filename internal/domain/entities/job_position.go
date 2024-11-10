package entities

import (
	"errors"
	"time"

	"github.com/Jollynjose/sistema-viatico-backend/internal/helpers"
	"github.com/google/uuid"
)

type JobPosition struct {
	ID                   uuid.UUID
	Name                 string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	JobPositionHistories []JobPositionHistory
}

func (j *JobPosition) validate() error {
	if j.Name == "" {
		return errors.New("name is required")
	}

	if !helpers.IsArrayEmpty(j.JobPositionHistories) {
		for _, jh := range j.JobPositionHistories {
			if err := jh.validate(); err != nil {
				return err
			}
		}
	}

	return nil
}

func NewJobPosition(name string, jobPositionHistories []JobPositionHistory) *JobPosition {
	jobPosition := &JobPosition{
		ID:                   uuid.New(),
		Name:                 name,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
		JobPositionHistories: jobPositionHistories,
	}

	return jobPosition
}
