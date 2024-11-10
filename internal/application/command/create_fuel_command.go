package command

import "github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"

type CreateFuelCommand struct {
	Type    db.FuelPriceType       `json:"type"`
	Name    string                 `json:"name"`
	History []CreatehistoryCommand `json:"fuel_histories"`
}
