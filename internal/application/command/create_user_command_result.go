package command

import "github.com/Jollynjose/sistema-viatico-backend/internal/application/common"

type CreateUserCommandResult struct {
	Result *common.UserResult `json:"result"`
}
