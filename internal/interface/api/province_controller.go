package api

import (
	"net/http"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/helpers"
)

type ProvinceController struct {
	service interfaces.ProvinceService
}

func NewProvinceController(router *http.ServeMux, service interfaces.ProvinceService) *ProvinceController {
	ctrl := &ProvinceController{
		service: service,
	}

	router.HandleFunc("POST /ingest", ctrl.IngestProvince)
	return ctrl
}

func (c *ProvinceController) IngestProvince(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.IngestProvince()

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
