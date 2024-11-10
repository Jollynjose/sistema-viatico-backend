package response

import (
	"time"

	"github.com/google/uuid"
)

type FindJobPositionResponse struct {
	ID                   uuid.UUID                        `json:"id"`
	Name                 string                           `json:"name"`
	CreatedAt            time.Time                        `json:"created_at"`
	UpdatedAt            time.Time                        `json:"updated_at"`
	JobPositionHistories []FindJobPositionHistoryResponse `json:"job_position_histories"`
}

type FindJobPositionsResponse struct {
	JobPositions []*FindJobPositionResponse `json:"job_positions"`
}
