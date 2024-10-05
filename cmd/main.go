package main

import (
	"fmt"
	"net/http"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/services"
	"github.com/Jollynjose/sistema-viatico-backend/internal/config"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	repository "github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db/repository/user"
	"github.com/Jollynjose/sistema-viatico-backend/internal/interface/api"
)

func main() {
	cfg := config.NewConfig()

	// Initialize the database
	gormDb := db.OpenDB(cfg)

	// Create Router
	mux := http.NewServeMux()

	// Repositories
	userRepository := repository.NewGormUserRepository(gormDb)

	// Services
	userService := services.NewUserService(userRepository)

	// Controllers
	api.NewUserController(mux, userService)

	// Create the server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.PORT),
		Handler: mux,
	}

	fmt.Println("Server is running on port ", cfg.PORT)
	// Initialize the server
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
	}

}
