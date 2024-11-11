package api

import (
	"net/http"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/helpers"
	"github.com/Jollynjose/sistema-viatico-backend/internal/interface/api/dto/request"
	"github.com/google/uuid"
)

type FuelController struct {
	service interfaces.FuelService
}

func NewFuelController(router *http.ServeMux, service interfaces.FuelService) *FuelController {
	ctrl := &FuelController{
		service: service,
	}

	router.HandleFunc("GET /", ctrl.GetAll)
	router.HandleFunc("GET /id", ctrl.GetByID)
	router.HandleFunc("POST /create", ctrl.Create)

	return ctrl
}

func (ctrl *FuelController) GetAll(w http.ResponseWriter, r *http.Request) {
	fuelsQuery, err := ctrl.service.FindAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	helpers.ResponseHandler(w, http.StatusOK, fuelsQuery)
}

func (ctrl *FuelController) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	fuelID, err := uuid.Parse(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fuelQuery, err := ctrl.service.FindByID(fuelID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	helpers.ResponseHandler(w, http.StatusOK, fuelQuery)
}

func (ctrl *FuelController) Create(w http.ResponseWriter, r *http.Request) {
	var createFuelRequest request.CreateFuelRequest

	err := helpers.DecodeJSONBody(w, r, &createFuelRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fuelCommand := createFuelRequest.ToCreateFuelCommand()

	fuel, err := ctrl.service.Create(fuelCommand)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.ResponseHandler(w, http.StatusCreated, fuel)
}
