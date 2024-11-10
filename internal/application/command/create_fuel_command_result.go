package command

import "github.com/Jollynjose/sistema-viatico-backend/internal/application/common"

type CreateFuelCommandResult struct {
	Result *common.FuelResult `json:"result"`
}
