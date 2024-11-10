package repositories

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/google/uuid"
)

type JobPositionRepository interface {
	Create(j *entities.JobPositionValidated) (*entities.JobPosition, error)
	FindById(id uuid.UUID) (*entities.JobPosition, error)
	FindAll() ([]*entities.JobPosition, error)
}
