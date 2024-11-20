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

	router.HandleFunc("GET /", ctrl.FindAll)

	return ctrl
}

func (c *ProvinceController) FindAll(w http.ResponseWriter, r *http.Request) {
	provinces, err := c.service.FindAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.ResponseHandler(w, http.StatusOK, provinces)
}
