package common

import (
	"time"

	"github.com/google/uuid"
)

type JobPositionResult struct {
	ID                   uuid.UUID                  `json:"id"`
	Name                 string                     `json:"name"`
	CreatedAt            time.Time                  `json:"created_at"`
	UpdatedAt            time.Time                  `json:"updated_at"`
	JobPositionHistories []JobPositionHistoryResult `json:"job_position_histories"`
}

func (j *JobPositionResult) GetMostRecentJobHistory() *JobPositionHistoryResult {
	if len(j.JobPositionHistories) == 0 {
		return nil
	}

	var mostRecent *JobPositionHistoryResult

	for _, history := range j.JobPositionHistories {
		if mostRecent == nil {
			mostRecent = &history
		} else if history.CreatedAt.After(mostRecent.CreatedAt) {
			mostRecent = &history
		}
	}

	return mostRecent
}
