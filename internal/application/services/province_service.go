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

type ProvinceService struct {
	ProvinceRepository repositories.ProvinceRepository
	cfg                *config.Config
}

func NewProvinceService(ProvinceRepository repositories.ProvinceRepository, cfg *config.Config) interfaces.ProvinceService {
	return &ProvinceService{
		ProvinceRepository: ProvinceRepository,
		cfg:                cfg,
	}
}

func (s *ProvinceService) IngestProvince() (*command.IngestProvinceCommandResult, error) {
	requestUrl := fmt.Sprintf("%s/%s", s.cfg.TERRITORIO_API_URL, "/provinces")

	res, err := http.Get(requestUrl)

	if err != nil {
		return nil, err
	}

	var provincesResponse struct {
		Valid bool `json:"valid"`
		Data  []struct {
			Name       string `json:"name"`
			Identifier string `json:"identifier"`
			Code       string `json:"code"`
			RegionCode string `json:"regionCode"`
		} `json:"data"`
	}

	err = json.NewDecoder(res.Body).Decode(&provincesResponse)

	if err != nil {
		return nil, err
	}

	var results []*common.ProvinceResult

	for _, province := range provincesResponse.Data {
		provinceEntity := entities.NewProvince(province.Name, province.Identifier, province.RegionCode, province.Code)

		validatedProvince := entities.NewProvinceValidated(provinceEntity)

		createdProvince, err := s.ProvinceRepository.Create(validatedProvince)

		if err != nil {
			return nil, err
		}

		result := mapper.NewProvinceResultFromValidatedEntity(createdProvince)

		results = append(results, result)
	}

	return &command.IngestProvinceCommandResult{
		Result: results,
	}, nil
}

func (s *ProvinceService) FindAll() (*query.FindAllProvinceQueryResult, error) {
	provinces, err := s.ProvinceRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var mappedProvinces []*common.ProvinceResult

	for _, province := range provinces {
		mappedProvince := mapper.NewProvinceResultFromValidatedEntity(province)

		mappedProvinces = append(mappedProvinces, mappedProvince)
	}

	return &query.FindAllProvinceQueryResult{
		Results: mappedProvinces,
	}, nil
}
