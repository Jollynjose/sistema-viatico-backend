package interfaces

import "github.com/Jollynjose/sistema-viatico-backend/internal/application/command"

type ProvinceService interface {
	IngestProvince() (*command.IngestProvinceCommandResult, error)
}
