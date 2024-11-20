package repositories

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/google/uuid"
)

type RouteRepository interface {
	FindAll() ([]*entities.Route, error)
	FindOne(id uuid.UUID) (*entities.Route, error)
	Create(routeValidated *entities.RouteValidated) (*entities.Route, error)
}
