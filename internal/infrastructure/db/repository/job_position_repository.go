package repository

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/repositories"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormJobPositionRepository struct {
	db *gorm.DB
}

func NewGormJobPositionRepository(db *gorm.DB) repositories.JobPositionRepository {
	return &GormJobPositionRepository{db: db}
}

func (r *GormJobPositionRepository) Create(j *entities.JobPositionValidated) (*entities.JobPosition, error) {
	dbJobPosition := toDbJobPosition(j)

	if err := r.db.Create(dbJobPosition).Error; err != nil {
		return nil, err
	}

	jobPosition := fromDbJobPosition(dbJobPosition)

	return jobPosition, nil
}

func (r *GormJobPositionRepository) FindById(id uuid.UUID) (*entities.JobPosition, error) {
	var dbJobPosition db.JobPosition

	err := r.db.Preload("JobPositionHistories").First(&dbJobPosition, "id = ?", id.String()).Error

	if err != nil {
		return nil, err
	}

	jobPosition := fromDbJobPosition(&dbJobPosition)

	return jobPosition, nil
}

func (r *GormJobPositionRepository) FindAll() ([]*entities.JobPosition, error) {
	var dbJobPositions []db.JobPosition

	err := r.db.Preload("JobPositionHistories").Find(&dbJobPositions).Error

	if err != nil {
		return nil, err
	}

	jobPositions := make([]*entities.JobPosition, len(dbJobPositions))

	for i, dbJobPosition := range dbJobPositions {
		jobPositions[i] = fromDbJobPosition(&dbJobPosition)
	}

	return jobPositions, nil
}
