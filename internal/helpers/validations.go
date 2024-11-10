package helpers

import (
	"fmt"

	"github.com/Jollynjose/sistema-viatico-backend/internal/config"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
)

func IsProduction(cfg *config.Config) bool {
	return cfg.ENV == config.PRODUCTION
}

func IsEmpty(s string) bool {
	return len(s) == 0
}

func IsAdmin(role string) bool {
	return db.Roles(role) == db.Admin_Role
}

func IsManager(role string) bool {
	return db.Roles(role) == db.Manager_Role
}

func IsGeneral(role string) bool {
	return db.Roles(role) == db.General_Role
}

func IsValidRole(role string) bool {
	fmt.Println(role)
	return IsAdmin(role) || IsManager(role) || IsGeneral(role)
}

func IsArrayEmpty[T any](arr []T) bool {
	return len(arr) == 0
}

func IsDiesel(fuelType string) bool {
	return db.FuelPriceType(fuelType) == db.Diesel_FuelPriceType
}

func IsGasoline(fuelType string) bool {
	return db.FuelPriceType(fuelType) == db.Gasoline_FuelPriceType
}

func IsValidFuelType(fuelType string) bool {
	return IsDiesel(fuelType) || IsGasoline(fuelType)
}
