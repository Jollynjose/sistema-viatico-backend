package api

import (
	"net/http"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/helpers"
)

type MunicipalityController struct {
	service interfaces.MunicipalityService
}

func NewMunicipalityController(router *http.ServeMux, service interfaces.MunicipalityService) *MunicipalityController {
	ctrl := &MunicipalityController{
		service: service,
	}

	router.HandleFunc("POST /ingest", ctrl.IngestMunicipality)
	return ctrl
}

func (c *MunicipalityController) IngestMunicipality(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.IngestMunicipality()

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
