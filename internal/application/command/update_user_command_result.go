package command

import "github.com/Jollynjose/sistema-viatico-backend/internal/application/common"

type UpdateUserCommandResult struct {
	Result *common.UserResult `json:"result"`
}
