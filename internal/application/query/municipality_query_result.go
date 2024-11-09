package query

import "github.com/Jollynjose/sistema-viatico-backend/internal/application/common"

type MunicipalityQueryResult struct {
	Result *common.MunicipalityResult `json:"result"`
}

type MunicipalitiesQueryResult struct {
	Results []*common.MunicipalityResult `json:"results"`
}
