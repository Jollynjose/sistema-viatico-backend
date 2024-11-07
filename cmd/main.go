package main

import (
	"fmt"
	"net/http"

	"github.com/Jollynjose/sistema-viatico-backend/internal/application/services"
	"github.com/Jollynjose/sistema-viatico-backend/internal/config"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db/repository"
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
	regionRouter := http.NewServeMux()
	provinceRouter := http.NewServeMux()
	municipalityRouter := http.NewServeMux()

	// Repositories
	userRepository := repository.NewGormUserRepository(gormDb)
	regionRepository := repository.NewGormRegionRepository(gormDb)
	provinceRepository := repository.NewGormProvinceRepository(gormDb)
	municipalityRepository := repository.NewGormMunicipalityRepository(gormDb)

	// Services
	userService := services.NewUserService(userRepository)
	regionService := services.NewRegionService(regionRepository, cfg)
	provinceService := services.NewProvinceService(provinceRepository, cfg)
	municipalityService := services.NewMunicipalityService(municipalityRepository, cfg)

	// Mount the userRouter
	mainRouter.Handle("/user/", http.StripPrefix("/user", middlewares.CheckAuthorization(userRouter)))
	mainRouter.Handle("/auth/", http.StripPrefix("/auth", authRouter))
	mainRouter.Handle("/region/", http.StripPrefix("/region", regionRouter))
	mainRouter.Handle("/province/", http.StripPrefix("/province", provinceRouter))
	mainRouter.Handle("/municipality/", http.StripPrefix("/municipality", municipalityRouter))

	// Controllers
	api.NewUserController(userRouter, userService)
	api.NewAuthController(authRouter, userService, cfg)
	api.NewRegionController(regionRouter, regionService)
	api.NewProvinceController(provinceRouter, provinceService)
	api.NewMunicipalityController(municipalityRouter, municipalityService)

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
