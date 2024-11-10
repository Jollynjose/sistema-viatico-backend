package repository

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/repositories"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FuelHistoryRepository struct {
	db *gorm.DB
}

func NewFuelHistoryRepository(db *gorm.DB) repositories.FuelHistoryRepository {
	return &FuelHistoryRepository{db}
}

func (r *FuelHistoryRepository) FindAllByFuelID(fuelID uuid.UUID) ([]*entities.FuelHistory, error) {
	var dbFuelHistories []db.FuelHistory

	if err := r.db.Find(&dbFuelHistories, "fuel_id = ?", fuelID.String()).Error; err != nil {
		return nil, err
	}

	var fuelHistories []*entities.FuelHistory

	for _, fh := range dbFuelHistories {
		fuelHistories = append(fuelHistories, &fromDBHistory(&fh).FuelHistory)
	}

	return fuelHistories, nil
}

func (r *FuelHistoryRepository) Create(f *entities.FuelHistoryValidated) (*entities.FuelHistory, error) {
	dbFuelHistory := toDBHistory(f)

	if err := r.db.Create(dbFuelHistory).Error; err != nil {
		return nil, err
	}

	fuelHistory := fromDBHistory(dbFuelHistory)

	return &fuelHistory.FuelHistory, nil
}

func (r *FuelHistoryRepository) FindCurrenthistory(fuelID uuid.UUID) (*entities.FuelHistory, error) {
	var dbFuelHistory db.FuelHistory

	err := r.db.
		First(&dbFuelHistory, "fuel_id = ?", fuelID.String()).
		Order("created_at DESC").
		Error

	if err != nil {
		return nil, err
	}

	fuelHistory := fromDBHistory(&dbFuelHistory)

	return &fuelHistory.FuelHistory, nil
}
