package response

import "time"

type FindUserResponse struct {
	ID            string    `json:"id"`
	Email         string    `json:"email"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Role          string    `json:"role"`
	JobPositionID string    `json:"job_position_id"`
}

type FindUsersResponse struct {
	Users []*FindUserResponse `json:"users"`
}
