package api

import (
	"net/http"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/helpers"
)

type RegionController struct {
	service interfaces.RegionService
}

func NewRegionController(router *http.ServeMux, service interfaces.RegionService) *RegionController {
	ctrl := &RegionController{
		service: service,
	}

	router.HandleFunc("POST /ingest", ctrl.IngestRegion)
	return ctrl
}

func (c *RegionController) IngestRegion(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.IngestRegion()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	if err := helpers.ParseJSON(w, result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
