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

	var usersResponse []*response.FindUserResponse

	for _, userQuery := range usersQuery.Results {
		usersResponse = append(usersResponse, mapper.ToFindUserResponse(userQuery))
	}

	response := response.FindUsersResponse{
		Users: usersResponse,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := helpers.ParseJSON(w, response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (uc *UserController) FindUserById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := helpers.ParseJSON(w, response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (uc *UserController) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := helpers.ParseJSON(w, response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
