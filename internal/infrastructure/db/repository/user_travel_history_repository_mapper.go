package repository

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
)

func ToDBUserTravelHistory(userTravel *entities.UserTravelHistory) *db.UserTravelHistory {
	return &db.UserTravelHistory{}
}

func FromDBUserTravelHistory(userTravel *db.UserTravelHistory) *entities.UserTravelHistory {
	return &entities.UserTravelHistory{}
}
