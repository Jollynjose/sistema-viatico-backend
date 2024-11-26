package repositories

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"github.com/google/uuid"
)

type TravelExpenseRepository interface {
	FindAll() ([]*entities.TravelExpense, error)
	Create(userTravel *entities.TravelExpense) (*entities.TravelExpense, error)
	FindOne(id uuid.UUID) (*db.TravelExpense, error)
}
