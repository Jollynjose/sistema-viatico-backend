package main

import (
	"fmt"
	"net/http"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/services"
	"github.com/Jollynjose/sistema-viatico-backend/internal/config"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	repository "github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db/repository/user"
	"github.com/Jollynjose/sistema-viatico-backend/internal/interface/api"
	"github.com/Jollynjose/sistema-viatico-backend/internal/interface/api/middlewares"
)

func main() {
	cfg := config.NewConfig()

	// Initialize the database
	gormDb := db.OpenDB(cfg)

	// Create Router
	mainRouter := http.NewServeMux()
	userRouter := http.NewServeMux()
	authRouter := http.NewServeMux()

	// Repositories
	userRepository := repository.NewGormUserRepository(gormDb)

	// Services
	userService := services.NewUserService(userRepository)

	// Mount the userRouter
	mainRouter.Handle("/user/", http.StripPrefix("/user", middlewares.CheckAuthorization(userRouter)))
	mainRouter.Handle("/auth/", http.StripPrefix("/auth", authRouter))

	// Controllers
	api.NewUserController(userRouter, userService)
	api.NewAuthController(authRouter, userService, cfg)

	// Create the server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.PORT),
		Handler: middlewares.Logger(mainRouter),
	}

	fmt.Println("Server is running on port ", cfg.PORT)
	// Initialize the server
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
	}

}
