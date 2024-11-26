package api

import (
	"net/http"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/helpers"
	"github.com/Jollynjose/sistema-viatico-backend/internal/interface/api/dto/request"
	"github.com/google/uuid"
)

type GeneratePdfController struct {
	travelExpenseService interfaces.TravelExpenseService
	fuelService          interfaces.FuelService
}

func NewGeneratePdfController(
	router *http.ServeMux,
	travelExpenseService interfaces.TravelExpenseService,
	fuelService interfaces.FuelService,
) *GeneratePdfController {
	controller := &GeneratePdfController{
		travelExpenseService: travelExpenseService,
		fuelService:          fuelService,
	}

	router.HandleFunc("POST /travel-expense", controller.GeneratePdf)

	return controller
}

func (gpc *GeneratePdfController) GeneratePdf(w http.ResponseWriter, r *http.Request) {
	var request request.GenerateTravelExpensePdf

	if err := helpers.DecodeJSONBody(w, r, &request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(request.TravelExpenseId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	travelExpense, err := gpc.travelExpenseService.FindOne(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fuel, err := gpc.fuelService.FindOneByFuelHistoryId(uuid.MustParse(travelExpense.FuelHistoryID))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate PDF
	pdf, err := helpers.GeneratePDF(travelExpense, fuel)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=travel_expense.pdf")
	w.Write(pdf)

}
