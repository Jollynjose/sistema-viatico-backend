package response

import "time"

type FindUserResponse struct {
	ID                      string    `json:"id"`
	Email                   string    `json:"email"`
	FirstName               string    `json:"first_name"`
	LastName                string    `json:"last_name"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
	Role                    string    `json:"role"`
	JobPositionID           string    `json:"job_position_id"`
	JobPostionSpecification *string   `json:"job_position_specification"`
}

type JobPosition struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Lunch         float64 `json:"lunch"`
	BreakFast     float64 `json:"breakfast"`
	Dinner        float64 `json:"dinner"`
	Accommodation float64 `json:"accommodation"`
}

type FindAllUser struct {
	FindUserResponse
	JobPosition JobPosition `json:"job_position"`
}

type FindUsersResponse struct {
	Users []*FindUserResponse `json:"users"`
}

type FindAllUsersResponse struct {
	Users []*FindAllUser `json:"users"`
}
