package api

import (
	"net/http"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/helpers"
	"github.com/Jollynjose/sistema-viatico-backend/internal/interface/api/dto/mapper"
	"github.com/Jollynjose/sistema-viatico-backend/internal/interface/api/dto/request"
	"github.com/Jollynjose/sistema-viatico-backend/internal/interface/api/dto/response"
	"github.com/google/uuid"
)

type JobPositionController struct {
	service interfaces.JobPositionService
}

func NewJobPositionController(router *http.ServeMux, service interfaces.JobPositionService) *JobPositionController {
	controller := &JobPositionController{
		service: service,
	}

	router.HandleFunc("GET /", controller.FindAll)
	router.HandleFunc("GET /{id}", controller.FindById)
	router.HandleFunc("POST /", controller.Create)
	return controller
}

func (ctrl *JobPositionController) FindAll(w http.ResponseWriter, r *http.Request) {
	jobPositionQuery, err := ctrl.service.FindAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var jobPositionResponse []*response.FindJobPositionResponse

	for _, jobPositionQuery := range jobPositionQuery.Results {
		jobPositionResponse = append(jobPositionResponse, mapper.ToJobPositionResponse(jobPositionQuery))
	}

	response := response.FindJobPositionsResponse{
		JobPositions: jobPositionResponse,
	}

	helpers.ResponseHandler(w, http.StatusOK, response)
}

func (ctrl *JobPositionController) FindById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	idToUuid, err := uuid.Parse(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jobPositionQuery, err := ctrl.service.FindById(&command.FindJobPositionByIdCommand{
		ID: idToUuid,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response := mapper.ToJobPositionResponse(jobPositionQuery.Result)

	helpers.ResponseHandler(w, http.StatusOK, response)
}

func (ctrl *JobPositionController) Create(w http.ResponseWriter, r *http.Request) {
	var createJobPositionRequest request.CreateJobPositionRequest

	err := helpers.DecodeJSONBody(w, r, &createJobPositionRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createJobPositionCommand := createJobPositionRequest.ToCreateJobPositionCommand()
	result, err := ctrl.service.Create(createJobPositionCommand)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := mapper.ToJobPositionResponse(result.Result)

	helpers.ResponseHandler(w, http.StatusCreated, response)
}
