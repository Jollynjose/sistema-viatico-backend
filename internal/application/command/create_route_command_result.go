package command

import "github.com/Jollynjose/sistema-viatico-backend/internal/application/common"

type CreateRouteCommandResult struct {
	Result *common.RouteResult `json:"result"`
}
