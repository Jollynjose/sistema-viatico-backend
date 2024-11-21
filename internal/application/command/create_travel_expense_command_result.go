package command

import "github.com/Jollynjose/sistema-viatico-backend/internal/application/common"

type CreateTravelExpenseCommandResult struct {
	Result *common.TravelExpenseResult `json:"result"`
}
