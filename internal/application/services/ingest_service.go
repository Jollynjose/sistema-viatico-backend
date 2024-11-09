package services

import (
	"errors"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/config"
)

type IngestionService struct {
	municipalityService interfaces.MunicipalityService
	provinceService     interfaces.ProvinceService
	regionService       interfaces.RegionService
	cfg                 *config.Config
}

func NewIngestionService(municipalityService interfaces.MunicipalityService, provinceService interfaces.ProvinceService, regionService interfaces.RegionService, cfg *config.Config) interfaces.IngestionService {
	return &IngestionService{
		municipalityService: municipalityService,
		provinceService:     provinceService,
		regionService:       regionService,
		cfg:                 cfg,
	}
}

func (s *IngestionService) IngestMapData() error {
	regions, err := s.IngestRegion()

	if err != nil {
		return err
	}

	if _, err = s.IngestProvince(); err != nil {
		return err
	}

	if _, err = s.IngestMunicipality(regions); err != nil {
		return err
	}

	return nil
}

func (s *IngestionService) IngestProvince() (*command.IngestProvinceCommandResult, error) {
	result, err := s.provinceService.IngestProvince()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *IngestionService) IngestRegion() (*command.IngestRegionCommandResult, error) {
	result, err := s.regionService.IngestRegion()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *IngestionService) IngestMunicipality(regions *command.IngestRegionCommandResult) (*command.IngestMunicipalityCommandResult, error) {
	var results []*common.MunicipalityResult

	for _, region := range regions.Result {
		ch := make(chan *command.IngestMunicipalityCommandResult)
		go func(region *common.RegionResult) {
			result, err := s.municipalityService.IngestMunicipality(region)

			if err != nil {
				ch <- nil
			}

			ch <- result
		}(region)

		result := <-ch

		if result == nil {
			return nil, errors.New("error ingesting municipalities")
		}

		results = append(results, result.Result...)
	}

	return &command.IngestMunicipalityCommandResult{
		Result: results,
	}, nil
}
