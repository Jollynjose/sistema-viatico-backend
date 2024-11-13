package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/mapper"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/query"
	"github.com/Jollynjose/sistema-viatico-backend/internal/config"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/repositories"
)

type RegionService struct {
	RegionRepository repositories.RegionRepository
	cfg              *config.Config
}

func NewRegionService(RegionRepository repositories.RegionRepository, cfg *config.Config) interfaces.RegionService {
	return &RegionService{
		RegionRepository: RegionRepository,
		cfg:              cfg,
	}
}

func (s *RegionService) IngestRegion() (*command.IngestRegionCommandResult, error) {
	requestUrl := fmt.Sprintf("%s/%s", s.cfg.TERRITORIO_API_URL, "/regions")

	res, err := http.Get(requestUrl)

	if err != nil {
		return nil, err
	}

	var regionsResponse struct {
		Valid bool `json:"valid"`
		Data  []struct {
			Name       string `json:"name"`
			Identifier string `json:"identifier"`
			Code       string `json:"code"`
		} `json:"data"`
	}

	err = json.NewDecoder(res.Body).Decode(&regionsResponse)

	if err != nil {
		return nil, err
	}

	var results []*common.RegionResult

	for _, region := range regionsResponse.Data {
		regionEntity := entities.NewRegion(region.Name, region.Identifier, region.Code)

		validatedRegion := entities.NewRegionValidated(regionEntity)

		createdRegion, err := s.RegionRepository.Create(validatedRegion)

		if err != nil {
			return nil, err
		}

		result := mapper.NewRegionResultFromValidatedEntity(createdRegion)

		results = append(results, result)
	}

	return &command.IngestRegionCommandResult{
		Result: results,
	}, nil
}

func (s *RegionService) FindAll() (*query.FindAllRegionQueryResult, error) {
	regions, err := s.RegionRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var results []*common.RegionResult

	for _, region := range regions {
		result := mapper.NewRegionResultFromValidatedEntity(region)

		results = append(results, result)
	}

	return &query.FindAllRegionQueryResult{
		Results: results,
	}, nil
}
