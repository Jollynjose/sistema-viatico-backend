package api

import (
	"net/http"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
)

type UserController struct {
	service interfaces.UserService
}

func NewUserController(router *http.ServeMux, service interfaces.UserService) *UserController {
	controller := &UserController{
		service: service,
	}

	router.HandleFunc("POST /hello", controller.hello)

	return controller
}

func (uc *UserController) hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
