package command

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
)

type IngestMunicipalityCommandResult struct {
	Result []*common.MunicipalityResult
}
