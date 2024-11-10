package query

import "github.com/Jollynjose/sistema-viatico-backend/internal/application/common"

type FuelHistoryQueryResult struct {
	Result *common.FuelHistoryResult `json:"result"`
}

type FuelHistoriesQueryResult struct {
	Results []*common.FuelHistoryResult `json:"results"`
}
