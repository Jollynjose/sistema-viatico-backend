package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/mapper"
	"github.com/Jollynjose/sistema-viatico-backend/internal/config"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/repositories"
)

type MunicipalityService struct {
	MunicipalityRepository repositories.MunicipalityRepository
	cfg                    *config.Config
}

func NewMunicipalityService(municipalityRepository repositories.MunicipalityRepository, cfg *config.Config) interfaces.MunicipalityService {
	return &MunicipalityService{
		MunicipalityRepository: municipalityRepository,
		cfg:                    cfg,
	}
}

func (s *MunicipalityService) IngestMunicipality() (*command.IngestMunicipalityCommandResult, error) {
	requestUrl := fmt.Sprintf("%s/%s", s.cfg.TERRITORIO_API_URL, "/municipalities")

	res, err := http.Get(requestUrl)

	if err != nil {
		return nil, err
	}

	var municipalitysResponse struct {
		Valid bool `json:"valid"`
		Data  []struct {
			Name       string `json:"name"`
			Identifier string `json:"identifier"`
			Code       string `json:"code"`
			RegionId   string `json:"regionCode"`
			ProvinceId string `json:"provinceCode"`
		} `json:"data"`
	}

	err = json.NewDecoder(res.Body).Decode(&municipalitysResponse)

	if err != nil {
		return nil, err
	}

	var results []*common.MunicipalityResult

	for _, municipality := range municipalitysResponse.Data {
		municipalityEntity := entities.NewMunicipality(municipality.Name, municipality.Identifier, municipality.RegionId, municipality.ProvinceId)

		validatedMunicipality := entities.NewMunicipalityValidated(municipalityEntity)

		createdMunicipality, err := s.MunicipalityRepository.Create(validatedMunicipality)

		if err != nil {
			return nil, err
		}

		result := mapper.NewMunicipalityResultFromValidatedEntity(createdMunicipality)

		results = append(results, result)
	}

	return &command.IngestMunicipalityCommandResult{
		Result: results,
	}, nil
}
