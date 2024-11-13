package command

import "github.com/Jollynjose/sistema-viatico-backend/internal/application/common"

type FindAllRegionCommandResult struct {
	Results []*common.RegionResult `json:"results"`
}
