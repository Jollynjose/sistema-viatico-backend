package api

import (
	"net/http"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/helpers"
)

type IngestionController struct {
	ingestionService interfaces.IngestionService
}

func NewIngestionController(router *http.ServeMux, ingestionService interfaces.IngestionService) *IngestionController {
	ctrl := &IngestionController{
		ingestionService: ingestionService,
	}

	router.HandleFunc("POST /data", ctrl.IngestMapData)

	return ctrl
}

func (c *IngestionController) IngestMapData(w http.ResponseWriter, r *http.Request) {
	err := c.ingestionService.IngestMapData()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	helpers.ResponseHandler(w, http.StatusCreated, []byte(`{"message": "Data Ingested"}`))
}
