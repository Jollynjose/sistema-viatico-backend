package api

import (
	"net/http"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/config"
	"github.com/Jollynjose/sistema-viatico-backend/internal/helpers"
	"github.com/Jollynjose/sistema-viatico-backend/internal/interface/api/dto/mapper"
	"github.com/Jollynjose/sistema-viatico-backend/internal/interface/api/dto/request"
)

type AuthController struct {
	userService interfaces.UserService
	cfg         *config.Config
}

func NewAuthController(router *http.ServeMux, userService interfaces.UserService, cfg *config.Config) *AuthController {
	controller := &AuthController{
		userService: userService,
		cfg:         cfg,
	}

	router.HandleFunc("POST /sign-up", controller.SignUp)
	router.HandleFunc("POST /sign-in", controller.SignIn)

	return controller
}

func (ac *AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	var signUpUserRequest request.SignUpUserRequest

	err := helpers.DecodeJSONBody(w, r, &signUpUserRequest)

	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	userCommand := signUpUserRequest.ToCreateUserCommand()

	user, err := ac.userService.Signup(userCommand)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create JWT token

	token, err := helpers.CreateToken(user.Result.ID.String(), ac.cfg.SECRET_KEY)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := mapper.ToAuthUserResponse(user.Result, token)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := helpers.ParseJSON(w, response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ac *AuthController) SignIn(w http.ResponseWriter, r *http.Request) {
	var signInUserRequest request.SignInUserRequest

	err := helpers.DecodeJSONBody(w, r, &signInUserRequest)

	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	userCommand := signInUserRequest.ToFindUserCommand()

	user, err := ac.userService.SignIn(userCommand)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create JWT token

	token, err := helpers.CreateToken(user.Result.ID.String(), ac.cfg.SECRET_KEY)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := mapper.ToAuthUserResponse(user.Result, token)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := helpers.ParseJSON(w, response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
