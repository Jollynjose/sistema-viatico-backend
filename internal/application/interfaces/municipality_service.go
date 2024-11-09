package interfaces

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/query"
)

type MunicipalityService interface {
	IngestMunicipality(region *common.RegionResult) (*command.IngestMunicipalityCommandResult, error)
	FindAll() (*query.MunicipalitiesQueryResult, error)
}
