package interfaces

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/query"
)

type JobPositionService interface {
	FindAll() (*query.JobPositionsQueryResult, error)
	FindById(jobPositionCommand *command.FindJobPositionByIdCommand) (*query.JobPositionQueryResult, error)
	Create(jobPositionCommand *command.CreateJobPositionCommand) (*command.CreateJobPositionCommandResult, error)
}
