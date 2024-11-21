package repositories

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/google/uuid"
)

type TravelExpenseRepository interface {
	FindAll() ([]*entities.TravelExpense, error)
	Create(userTravel *entities.TravelExpense) (*entities.TravelExpense, error)
	FindOne(id uuid.UUID) (*entities.TravelExpense, error)
}
