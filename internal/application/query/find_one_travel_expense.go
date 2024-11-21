package query

import "github.com/Jollynjose/sistema-viatico-backend/internal/application/common"

type FindOneTravelExpenseQuery struct {
	Result *common.TravelExpenseResult `json:"result"`
}
