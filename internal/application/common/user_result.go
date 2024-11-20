package common

import (
	"time"

	"github.com/google/uuid"
)

type UserResult struct {
	ID            uuid.UUID          `json:"id"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
	FirstName     string             `json:"first_name"`
	LastName      string             `json:"last_name"`
	Email         string             `json:"email"`
	Role          string             `json:"role"`
	JobPositionID string             `json:"job_position_id"`
	JobPosition   *JobPositionResult `json:"job_position"`
}
