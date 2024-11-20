package repositories

import "github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"

type JobPositionHistoryRepository interface {
	Create(jph *entities.JobPositionHistoryValidated) (*entities.JobPositionHistory, error)
}
