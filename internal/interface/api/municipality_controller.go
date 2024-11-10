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
