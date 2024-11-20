package query

import "github.com/Jollynjose/sistema-viatico-backend/internal/application/common"

type FindRouteQueryResult struct {
	Result *common.RouteResult `json:"result"`
}
