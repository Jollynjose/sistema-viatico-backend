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

	router.HandleFunc("GET /", ctrl.FindAll)

	router.HandleFunc("GET /by_region_code/{id}", ctrl.FindByRegionCode)
	return ctrl
}

func (c *MunicipalityController) FindAll(w http.ResponseWriter, r *http.Request) {
	municipalitiesQuery, err := c.service.FindAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.ResponseHandler(w, http.StatusOK, municipalitiesQuery)
}

func (c *MunicipalityController) FindByRegionCode(w http.ResponseWriter, r *http.Request) {
	regionCode := r.PathValue("id")

	if helpers.IsEmpty(regionCode) {
		http.Error(w, "Region code is required", http.StatusBadRequest)
		return
	}

	municipalitiesQuery, err := c.service.FindByRegionCode(regionCode)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.ResponseHandler(w, http.StatusOK, municipalitiesQuery)
}
