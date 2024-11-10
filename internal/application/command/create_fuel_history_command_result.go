package command

import "github.com/Jollynjose/sistema-viatico-backend/internal/application/common"

type CreatehistoryCommandResults struct {
	Result *common.FuelResult `json:"result"`
}
