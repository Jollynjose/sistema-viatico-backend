package interfaces

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
)

type JobPositionHistoriesService interface {
	CreateJobPositionHistory(jph *command.CreateJobPositionHistoryCommandByJobPositionID) (*command.CreateJobPositionHistoryCommandResult, error)
}
