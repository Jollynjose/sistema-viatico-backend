package repository

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/repositories"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RouteRepository struct {
	db *gorm.DB
}

func NewGormRouteRepository(db *gorm.DB) repositories.RouteRepository {
	return &RouteRepository{
		db: db,
	}
}

func (r *RouteRepository) FindAll() ([]*entities.Route, error) {
	var dbRoutes []db.Route

	if err := r.db.Find(&dbRoutes).Error; err != nil {
		return nil, err
	}

	var routes []*entities.Route

	for _, dbRoute := range dbRoutes {
		routes = append(routes, FromDBRoute(&dbRoute))
	}

	return routes, nil
}

func (r *RouteRepository) FindOne(id uuid.UUID) (*entities.Route, error) {
	var route db.Route

	if err := r.db.First(&route, id).Error; err != nil {
		return nil, err
	}

	return FromDBRoute(&route), nil
}

func (r *RouteRepository) Create(routeValidated *entities.RouteValidated) (*entities.Route, error) {
	route := ToDBRoute(routeValidated)

	if err := r.db.Create(route).Error; err != nil {
		return nil, err
	}

	return FromDBRoute(route), nil
}
