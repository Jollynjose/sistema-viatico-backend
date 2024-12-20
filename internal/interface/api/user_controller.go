package api

import (
	"net/http"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/command"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/common"
	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/helpers"
	"github.com/Jollynjose/sistema-viatico-backend/internal/interface/api/dto/mapper"
	"github.com/Jollynjose/sistema-viatico-backend/internal/interface/api/dto/request"
	"github.com/Jollynjose/sistema-viatico-backend/internal/interface/api/dto/response"
	"github.com/google/uuid"
)

type UserController struct {
	service interfaces.UserService
}

func NewUserController(router *http.ServeMux, service interfaces.UserService) *UserController {
	controller := &UserController{
		service: service,
	}

	router.HandleFunc("POST /hello", controller.hello)
	router.HandleFunc("GET /", controller.FindAll)
	router.HandleFunc("GET /{id}", controller.FindUserById)
	router.HandleFunc("PUT /{id}", controller.UpdateUserById)

	return controller
}

func (uc *UserController) hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func (uc *UserController) FindAll(w http.ResponseWriter, r *http.Request) {

	usersQuery, err := uc.service.FindAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var usersResponse []*response.FindAllUser

	for _, userQuery := range usersQuery.Results {
		jobPosition := userQuery.JobPosition
		var jobPositionHistory *common.JobPositionHistoryResult

		if jobPosition != nil {
			jobPositionHistory = jobPosition.GetMostRecentJobHistory()
		}

		var jp response.JobPosition

		if jobPositionHistory != nil {
			jp = response.JobPosition{
				ID:                   jobPosition.ID.String(),
				Name:                 jobPosition.Name,
				Lunch:                jobPositionHistory.Lunch,
				BreakFast:            jobPositionHistory.BreakFast,
				Dinner:               jobPositionHistory.Dinner,
				Accommodation:        jobPositionHistory.Accommodation,
				JobPositionHistoryId: jobPositionHistory.ID.String(),
			}
		}

		usersResponse = append(usersResponse, mapper.ToFindAllUser(userQuery, &jp))
	}

	response := response.FindAllUsersResponse{
		Users: usersResponse,
	}

	helpers.ResponseHandler(w, http.StatusOK, response)
}

func (uc *UserController) FindUserById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := uuid.Validate(id)

	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	userQuery, err := uc.service.FindUserById(&command.FindUserByIdCommand{
		ID: id,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := mapper.ToFindUserResponse(userQuery.Result)

	helpers.ResponseHandler(w, http.StatusOK, response)
}

func (uc *UserController) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := uuid.Validate(id)

	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	var updateUserRequest request.UpdateUserByIdRequest

	err = helpers.DecodeJSONBody(w, r, &updateUserRequest)

	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	userCommand := updateUserRequest.ToUpdateUserCommand()

	user, err := uc.service.UpdateById(id, userCommand)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := mapper.ToFindUserResponse(user.Result)

	helpers.ResponseHandler(w, http.StatusCreated, response)
}
