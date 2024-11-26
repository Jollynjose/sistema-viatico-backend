package repository

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/repositories"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TravelExpenseRepository struct {
	db *gorm.DB
}

func NewGormTravelExpenseRepository(db *gorm.DB) repositories.TravelExpenseRepository {
	return &TravelExpenseRepository{
		db: db,
	}
}

func (r *TravelExpenseRepository) FindAll() ([]*entities.TravelExpense, error) {
	var userTravels []db.TravelExpense

	if err := r.db.Find(&userTravels).Error; err != nil {
		return nil, err
	}

	var userTravelEntities []*entities.TravelExpense

	for _, userTravel := range userTravels {
		userTravelEntities = append(userTravelEntities, FromDBUserTravel(&userTravel))
	}

	return userTravelEntities, nil
}

func (r *TravelExpenseRepository) Create(userTravel *entities.TravelExpense) (*entities.TravelExpense, error) {
	userTravelDB := ToDBUserTravel(userTravel)

	if err := r.db.Create(userTravelDB).Error; err != nil {
		return nil, err
	}

	return FromDBUserTravel(userTravelDB), nil
}

func (r *TravelExpenseRepository) FindOne(id uuid.UUID) (*db.TravelExpense, error) {
	var travelExpense db.TravelExpense
	if err := r.db.
		Preload("UserTravelHistory.User.JobPosition").
		Preload("UserTravelHistory.JobPositionHistory").
		Preload("Route.StartingPointProvince").
		Preload("Route.FinalDestinationProvince").
		Preload(clause.Associations).
		Where("id = ?", id).
		First(&travelExpense).Error; err != nil {
		return nil, err
	}

	return &travelExpense, nil
}
