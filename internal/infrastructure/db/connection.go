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

	db.AutoMigrate(&User{}, &Region{}, &Province{}, &Municipality{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
