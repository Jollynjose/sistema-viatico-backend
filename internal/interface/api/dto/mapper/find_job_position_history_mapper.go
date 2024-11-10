package mapper

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/interface/api/dto/response"
)

func ToJobPositionHistoryResponse(j *common.JobPositionHistoryResult) *response.FindJobPositionHistoryResponse {
	return &response.FindJobPositionHistoryResponse{
		ID:            j.ID,
		JobPositionID: j.JobPositionID,
		Lunch:         j.Lunch,
		BreakFast:     j.BreakFast,
		Dinner:        j.Dinner,
		Accommodation: j.Accommodation,
		CreatedAt:     j.CreatedAt,
		UpdatedAt:     j.UpdatedAt,
	}
}
