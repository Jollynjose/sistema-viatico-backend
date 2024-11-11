package request

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
)

type CreateJobPositionHistoriesRequest struct {
	Lunch         float64 `json:"lunch"`
	BreakFast     float64 `json:"breakfast"`
	Dinner        float64 `json:"dinner"`
	Accommodation float64 `json:"accommodation"`
}

type CreateJobPositionRequest struct {
	Name                 string                              `json:"name"`
	JobPositionHistories []CreateJobPositionHistoriesRequest `json:"job_position_histories"`
}

func (r *CreateJobPositionRequest) ToCreateJobPositionCommand() *command.CreateJobPositionCommand {
	var jobPositionHistories []command.CreateJobPositionHistoryCommand

	for _, history := range r.JobPositionHistories {
		jobPositionHistories = append(jobPositionHistories, command.CreateJobPositionHistoryCommand{
			Lunch:         history.Lunch,
			BreakFast:     history.BreakFast,
			Dinner:        history.Dinner,
			Accommodation: history.Accommodation,
		})
	}

	return &command.CreateJobPositionCommand{
		Name:                 r.Name,
		JobPositionHistories: jobPositionHistories,
	}
}
