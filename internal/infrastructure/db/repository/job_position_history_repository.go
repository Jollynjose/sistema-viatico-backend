package repository

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/repositories"
	"gorm.io/gorm"
)

type JobPositionHistoryRepository struct {
	db *gorm.DB
}

func NewGormJobPositionHistoryRepository(db *gorm.DB) repositories.JobPositionHistoryRepository {
	return &JobPositionHistoryRepository{
		db: db,
	}
}

func (jphr *JobPositionHistoryRepository) Create(jph *entities.JobPositionHistoryValidated) (*entities.JobPositionHistory, error) {
	jobPositionDb := toDBJobPositionHistory(&jph.JobPositionHistory)

	if err := jphr.db.Create(&jobPositionDb).Error; err != nil {
		return nil, err
	}

	return fromDBJobPositionHistory(jobPositionDb), nil
}
