package query

import "github.com/Jollynjose/sistema-viatico-backend/internal/application/common"

type FindAllRegionQueryResult struct {
	Results []*common.RegionResult `json:"results"`
}
