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

	router.HandleFunc("GET /a", ctrl.FindAll)
	return ctrl
}

func (ctrl *RegionController) FindAll(w http.ResponseWriter, r *http.Request) {
	regions, err := ctrl.service.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.ResponseHandler(w, http.StatusOK, regions)
}
