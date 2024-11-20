package query

import "github.com/Jollynjose/sistema-viatico-backend/internal/application/common"

type FindAllRouteQueryResult struct {
	Results []*common.RouteResult `json:"results"`
}
