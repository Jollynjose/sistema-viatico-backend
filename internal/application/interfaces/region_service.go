package interfaces

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/query"
)

type RegionService interface {
	IngestRegion() (*command.IngestRegionCommandResult, error)
	FindAll() (*query.FindAllRegionQueryResult, error)
}
