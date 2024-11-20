package query

import "github.com/Jollynjose/sistema-viatico-backend/internal/application/common"

type FindAllProvinceQueryResult struct {
	Results []*common.ProvinceResult `json:"results"`
}
