package command

import "github.com/Jollynjose/sistema-viatico-backend/internal/application/common"

type CreateJobPositionHistoryCommandResult struct {
	Result *common.JobPositionHistoryResult `json:"result"`
}
