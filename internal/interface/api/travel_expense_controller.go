package api

import (
	"net/http"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/helpers"
	"github.com/Jollynjose/sistema-viatico-backend/internal/interface/api/dto/request"
	"github.com/google/uuid"
)

type TravelExpenseController struct {
	service interfaces.TravelExpenseService
}

func NewTravelExpenseController(router *http.ServeMux, service interfaces.TravelExpenseService) *TravelExpenseController {
	ctrl := &TravelExpenseController{service: service}

	router.HandleFunc("POST /create", ctrl.Create)
	router.HandleFunc("GET /{id}", ctrl.FindOne)
	router.HandleFunc("GET /pdf", ctrl.PDF)

	return ctrl
}

func (ctrl *TravelExpenseController) Create(w http.ResponseWriter, r *http.Request) {
	var travelExpenseRequest request.CreateTravelExpenseRequest

	err := helpers.DecodeJSONBody(w, r, &travelExpenseRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	travelExpenseCommand, err := travelExpenseRequest.ToCreateTravelExpenseCommand()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	travelExpenseResult, err := ctrl.service.Create(travelExpenseCommand)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.ResponseHandler(w, http.StatusCreated, travelExpenseResult)
}

func (ctrl *TravelExpenseController) FindOne(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	idToUUID, err := uuid.Parse(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	travelExpense, err := ctrl.service.FindOne(idToUUID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.ResponseHandler(w, http.StatusOK, travelExpense)
}

func (ctrl *TravelExpenseController) PDF(w http.ResponseWriter, r *http.Request) {
	// pdf, err := helpers.GeneratePDF()

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	w.Header().Set("Content-Type", "application/pdf")
	// w.Write(pdf)
	w.WriteHeader(http.StatusOK)
}
