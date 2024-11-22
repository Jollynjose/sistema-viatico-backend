package db

import (
	"fmt"

	"github.com/Jollynjose/sistema-viatico-backend/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DB_HOST,
		cfg.DB_USER,
		cfg.DB_PASSWORD,
		cfg.DB_NAME,
		cfg.DB_PORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	if cfg.ENV == "dev" {
		db.Exec(`
		DO $$
		Begin
			IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'roles') THEN
				RAISE NOTICE 'roles ENUM already exists';
			ELSE
				CREATE TYPE roles AS ENUM ('admin', 'user');
			END IF;
			
			IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'fuel_price_type') THEN
				RAISE NOTICE 'fuel_price_type ENUM already exists';
			ELSE
				CREATE TYPE fuel_price_type AS ENUM ('gasoline', 'diesel');
			END IF;
		End;
		$$
	`)

		db.AutoMigrate(
			&JobPosition{},
			&JobPositionHistory{},
			&Region{},
			&Province{},
			&Municipality{},
			&Route{},
			&Toll{},
			&Fuel{},
			&FuelHistory{},
			&UserTravelHistory{},
			&TravelExpense{},
			&User{},
		)
	}

	// Uncomment the following lines to create ENUM types in the database

	return db
}
