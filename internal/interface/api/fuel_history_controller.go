package api

import (
	"net/http"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/helpers"
	"github.com/Jollynjose/sistema-viatico-backend/internal/interface/api/dto/request"
)

type FuelHistoryController struct {
	service interfaces.FuelHistoryService
}

func NewFuelHistoryController(router *http.ServeMux, service interfaces.FuelHistoryService) *FuelHistoryController {
	ctrl := &FuelHistoryController{
		service: service,
	}

	router.HandleFunc("POST /create", ctrl.Create)

	return ctrl
}

func (ctrl *FuelHistoryController) Create(w http.ResponseWriter, r *http.Request) {
	var createFuelHistoryRequest request.CreateFuelHistoryRequest

	err := helpers.DecodeJSONBody(w, r, &createFuelHistoryRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createFuelHistoryCommand := createFuelHistoryRequest.ToCreateFuelHistoryCommand()
	result, err := ctrl.service.Create(createFuelHistoryCommand)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.ResponseHandler(w, http.StatusCreated, result)
}
