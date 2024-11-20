package services

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/mapper"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/query"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/repositories"
	"github.com/google/uuid"
)

type RouteService struct {
	repo repositories.RouteRepository
}

func NewRouteService(repo repositories.RouteRepository) interfaces.RouteService {
	return &RouteService{repo: repo}
}

func (s *RouteService) FindAll() (*query.FindAllRouteQueryResult, error) {
	routes, err := s.repo.FindAll()

	if err != nil {
		return nil, err
	}

	var results []*common.RouteResult

	for _, route := range routes {
		results = append(results, mapper.NewRouteResultFromEntity(route))
	}

	return &query.FindAllRouteQueryResult{
		Results: results,
	}, nil
}

func (s *RouteService) FindById(id uuid.UUID) (*query.FindRouteQueryResult, error) {

	route, err := s.repo.FindOne(id)

	if err != nil {
		return nil, err
	}

	result := mapper.NewRouteResultFromEntity(route)

	return &query.FindRouteQueryResult{
		Result: result,
	}, nil
}

func (s *RouteService) Create(route *command.CreateRouteCommand) (*command.CreateRouteCommandResult, error) {
	newRoute := entities.NewRoute(
		route.StartingPointProvinceID,
		route.FinalDestinationProvinceID,
		route.Description,
		route.TotalKms,
	)

	validatedRoute := entities.NewRouteValidated(*newRoute)

	if validatedRoute.Err != nil {
		return nil, validatedRoute.Err
	}

	createdRoute, err := s.repo.Create(validatedRoute)

	if err != nil {
		return nil, err
	}

	return &command.CreateRouteCommandResult{
		Result: mapper.NewRouteResultFromEntity(createdRoute),
	}, nil
}
