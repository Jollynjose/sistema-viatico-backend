package interfaces

import "github.com/Jollynjose/sistema-viatico-backend/internal/application/command"

type RegionService interface {
	IngestRegion() (*command.IngestRegionCommandResult, error)
}
