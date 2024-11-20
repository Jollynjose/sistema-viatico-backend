package command

import "github.com/google/uuid"

type CreateRouteCommand struct {
	StartingPointProvinceID    uuid.UUID `json:"starting_point_province_id"`
	FinalDestinationProvinceID uuid.UUID `json:"final_destination_province_id"`
	Description                string    `json:"description"`
	TotalKms                   int       `json:"total_kms"`
}
