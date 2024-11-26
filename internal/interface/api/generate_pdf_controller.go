package api

import (
	"net/http"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/helpers"
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

	router.HandleFunc("GET /travel-expense/{id}", controller.GeneratePdf)

	return controller
}

func (gpc *GeneratePdfController) GeneratePdf(w http.ResponseWriter, r *http.Request) {
	idRaw := r.PathValue("id")

	id, err := uuid.Parse(idRaw)

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
