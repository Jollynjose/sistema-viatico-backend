package mapper

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/interface/api/dto/response"
)

func ToJobPositionResponse(j *common.JobPositionResult) *response.FindJobPositionResponse {
	var jobPositionHistoryResponses []response.FindJobPositionHistoryResponse

	for _, jph := range j.JobPositionHistories {
		jobPositionHistoryResponses = append(jobPositionHistoryResponses, *ToJobPositionHistoryResponse(&jph))
	}

	return &response.FindJobPositionResponse{
		ID:                   j.ID,
		Name:                 j.Name,
		CreatedAt:            j.CreatedAt,
		UpdatedAt:            j.UpdatedAt,
		JobPositionHistories: jobPositionHistoryResponses,
	}
}
