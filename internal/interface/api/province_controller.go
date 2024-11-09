package api

import (
	"net/http"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
)

type ProvinceController struct {
	service interfaces.ProvinceService
}

func NewProvinceController(router *http.ServeMux, service interfaces.ProvinceService) *ProvinceController {
	ctrl := &ProvinceController{
		service: service,
	}

	return ctrl
}
