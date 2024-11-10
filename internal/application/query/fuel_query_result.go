package query

import "github.com/Jollynjose/sistema-viatico-backend/internal/application/common"

type FuelQueryResult struct {
	Result *common.FuelResult `json:"result"`
}

type FuelsQueryResult struct {
	Results []*common.FuelResult `json:"results"`
}
