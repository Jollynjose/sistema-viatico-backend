package interfaces

import "github.com/Jollynjose/sistema-viatico-backend/internal/application/command"

type MunicipalityService interface {
	IngestMunicipality() (*command.IngestMunicipalityCommandResult, error)
}
