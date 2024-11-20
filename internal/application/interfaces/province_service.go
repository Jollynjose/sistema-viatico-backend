package interfaces

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/query"
)

type ProvinceService interface {
	IngestProvince() (*command.IngestProvinceCommandResult, error)
	FindAll() (*query.FindAllProvinceQueryResult, error)
}
