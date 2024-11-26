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

func NewGormFuelRepository(db *gorm.DB) repositories.FuelRepository {
	return &FuelRepository{db}
}

func (r *FuelRepository) FindAll() ([]*entities.Fuel, error) {
	var dbFuels []db.Fuel

	if err := r.db.Preload("History").Find(&dbFuels).Error; err != nil {
		return nil, err
	}

	var fuels []*entities.Fuel

	for _, f := range dbFuels {
		fuel := fromDBFuel(&f)
		fuels = append(fuels, fuel)
	}

	return fuels, nil
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

func (r *FuelRepository) FindOneByFuelHistoryId(fuelHistoryId uuid.UUID) (*db.Fuel, error) {
	var dbFuel db.Fuel

	err := r.db.Joins("History", r.db.Where(&db.FuelHistory{FuelID: fuelHistoryId.String()})).First(&dbFuel)

	if err.Error != nil {
		return nil, err.Error
	}

	return &dbFuel, nil
}
