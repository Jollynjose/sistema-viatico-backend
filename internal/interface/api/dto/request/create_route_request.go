package request

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/google/uuid"
)

type CreateRouteRequest struct {
	StartingPointProvinceID    string `json:"starting_point_province_id"`
	FinalDestinationProvinceID string `json:"final_destination_province_id"`
	Description                string `json:"description"`
	TotalKms                   int    `json:"total_kms"`
}

func (r *CreateRouteRequest) ToCreateRouteCommand() *command.CreateRouteCommand {
	return &command.CreateRouteCommand{
		StartingPointProvinceID:    uuid.MustParse(r.StartingPointProvinceID),
		FinalDestinationProvinceID: uuid.MustParse(r.FinalDestinationProvinceID),
		Description:                r.Description,
		TotalKms:                   r.TotalKms,
	}
}
