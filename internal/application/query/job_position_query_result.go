package query

import "github.com/Jollynjose/sistema-viatico-backend/internal/application/common"

type JobPositionQueryResult struct {
	Result *common.JobPositionResult `json:"result"`
}

type JobPositionsQueryResult struct {
	Results []*common.JobPositionResult `json:"results"`
}
