package command

import "github.com/google/uuid"

type CreateJobPositionHistoryCommand struct {
	Lunch         float64 `json:"lunch"`
	BreakFast     float64 `json:"breakfast"`
	Dinner        float64 `json:"dinner"`
	Accommodation float64 `json:"accommodation"`
}

type CreateJobPositionHistoryCommandByJobPositionID struct {
	CreateJobPositionHistoryCommand
	JobPositionID uuid.UUID `json:"job_position_id"`
}
