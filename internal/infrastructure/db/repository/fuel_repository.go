package repository

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/repositories"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FuelRepository struct {
	db *gorm.DB
}

func NewFuelRepository(db *gorm.DB) repositories.FuelRepository {
	return &FuelRepository{db}
}

func (r *FuelRepository) FindAll() ([]*entities.Fuel, error) {
	var dbFuels []db.Fuel

	if err := r.db.Preload("History").Find(&dbFuels).Error; err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *FuelRepository) Create(f *entities.FuelValidated) (*entities.Fuel, error) {
	dbFuel := toDBFuel(f)

	if err := r.db.Create(dbFuel).Error; err != nil {
		return nil, err
	}

	fuel := fromDBFuel(dbFuel)

	return fuel, nil
}

func (r *FuelRepository) FindByID(ID uuid.UUID) (*entities.Fuel, error) {
	var dbFuel db.Fuel

	if err := r.db.Preload("History").First(&dbFuel, "id = ?", ID.String()).Error; err != nil {
		return nil, err
	}

	fuel := fromDBFuel(&dbFuel)

	return fuel, nil
}
