package api

import (
	"net/http"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/helpers"
	"github.com/Jollynjose/sistema-viatico-backend/internal/interface/api/dto/request"
)

type JobPositionHistoryController struct {
	service interfaces.JobPositionHistoriesService
}

func NewJobPositionHistoryController(router *http.ServeMux, service interfaces.JobPositionHistoriesService) *JobPositionHistoryController {
	controller := &JobPositionHistoryController{
		service: service,
	}

	router.HandleFunc("/job_position_history", controller.CreateJobPositionHistory)

	return controller
}

func (jphc *JobPositionHistoryController) CreateJobPositionHistory(w http.ResponseWriter, r *http.Request) {
	var jobPositionHistoryRequest request.CreateJobPositionHistoryRequest

	if err := helpers.DecodeJSONBody(w, r, &jobPositionHistoryRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jobPositionHistoryCommand := jobPositionHistoryRequest.ToCreateJobPositionHistoryCommand()

	jobPositionHistoryResult, err := jphc.service.CreateJobPositionHistory(jobPositionHistoryCommand)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.ResponseHandler(w, http.StatusCreated, jobPositionHistoryResult)
}
