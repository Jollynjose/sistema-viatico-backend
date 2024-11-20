package request

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/google/uuid"
)

type CreateJobPositionHistoryRequest struct {
	JobPositionID uuid.UUID `json:"job_position_id"`
	Lunch         float64   `json:"lunch"`
	BreakFast     float64   `json:"breakfast"`
	Dinner        float64   `json:"dinner"`
	Accommodation float64   `json:"accommodation"`
}

func (c *CreateJobPositionHistoryRequest) ToCreateJobPositionHistoryCommand() *command.CreateJobPositionHistoryCommandByJobPositionID {
	return &command.CreateJobPositionHistoryCommandByJobPositionID{
		CreateJobPositionHistoryCommand: command.CreateJobPositionHistoryCommand{
			Lunch:         c.Lunch,
			BreakFast:     c.BreakFast,
			Dinner:        c.Dinner,
			Accommodation: c.Accommodation,
		},
		JobPositionID: c.JobPositionID,
	}
}
