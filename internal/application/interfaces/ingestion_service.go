package interfaces

import "github.com/Jollynjose/sistema-viatico-backend/internal/application/command"

type IngestionService interface {
	IngestMapData() error
	IngestProvince() (*command.IngestProvinceCommandResult, error)
	IngestRegion() (*command.IngestRegionCommandResult, error)
	IngestMunicipality(regions *command.IngestRegionCommandResult) (*command.IngestMunicipalityCommandResult, error)
}
