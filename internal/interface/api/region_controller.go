package api

import (
	"net/http"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
)

type RegionController struct {
	service interfaces.RegionService
}

func NewRegionController(router *http.ServeMux, service interfaces.RegionService) *RegionController {
	ctrl := &RegionController{
		service: service,
	}

	return ctrl
}
