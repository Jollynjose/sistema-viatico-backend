package api

import (
	"net/http"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/interfaces"
	"github.com/Jollynjose/sistema-viatico-backend/internal/helpers"
	"github.com/Jollynjose/sistema-viatico-backend/internal/interface/api/dto/request"
	"github.com/google/uuid"
)

type RouteController struct {
	service interfaces.RouteService
}

func NewRouteController(router *http.ServeMux, service interfaces.RouteService) *RouteController {
	ctrl := &RouteController{
		service: service,
	}

	router.HandleFunc("GET /", ctrl.GetRoutes)
	router.HandleFunc("GET /:id", ctrl.GetRoute)
	router.HandleFunc("POST /p", ctrl.CreateRoute)

	return ctrl
}

func (ctrl *RouteController) GetRoutes(w http.ResponseWriter, r *http.Request) {
	routes, err := ctrl.service.FindAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.ResponseHandler(w, http.StatusOK, routes)
}

func (ctrl *RouteController) GetRoute(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	idToUUID, err := uuid.Parse(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	route, err := ctrl.service.FindById(idToUUID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.ResponseHandler(w, http.StatusOK, route)
}

func (ctrl *RouteController) CreateRoute(w http.ResponseWriter, r *http.Request) {
	var routeRequest request.CreateRouteRequest

	err := helpers.DecodeJSONBody(w, r, &routeRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	routeToCommand := routeRequest.ToCreateRouteCommand()

	route, err := ctrl.service.Create(routeToCommand)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.ResponseHandler(w, http.StatusCreated, route)
}
