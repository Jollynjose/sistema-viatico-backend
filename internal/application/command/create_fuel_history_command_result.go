package command

import "github.com/Jollynjose/sistema-viatico-backend/internal/application/common"

type CreateFuelHistoryCommandResults struct {
	Result *common.FuelHistoryResult `json:"result"`
}
