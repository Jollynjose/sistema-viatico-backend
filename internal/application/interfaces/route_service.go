package interfaces

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/query"
	"github.com/google/uuid"
)

type RouteService interface {
	FindAll() (*query.FindAllRouteQueryResult, error)
	FindById(id uuid.UUID) (*query.FindRouteQueryResult, error)
	Create(route *command.CreateRouteCommand) (*command.CreateRouteCommandResult, error)
}
