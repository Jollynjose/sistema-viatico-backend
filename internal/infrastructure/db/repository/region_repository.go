package repository

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/repositories"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"gorm.io/gorm"
)

type GormRegionRepository struct {
	db *gorm.DB
}

func NewGormRegionRepository(db *gorm.DB) repositories.RegionRepository {
	return &GormRegionRepository{db: db}
}

func (r *GormRegionRepository) Create(user *entities.RegionValidated) (*entities.Region, error) {
	dbData := toDBRegion(user)

	if err := r.db.Create(&dbData).Error; err != nil {
		return nil, err
	}

	return fromDBRegion(dbData), nil
}

func (r *GormRegionRepository) FindById(id string) (*entities.Region, error) {
	var dbRegion db.Region

	err := r.db.First(&dbRegion, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return fromDBRegion(&dbRegion), nil
}

func (r *GormRegionRepository) FindByIdentifier(identifier string) (*entities.Region, error) {
	var dbRegion db.Region

	err := r.db.First(&dbRegion, "identifier = ?", identifier).Error

	if err != nil {
		return nil, err
	}

	return fromDBRegion(&dbRegion), nil
}

func (r *GormRegionRepository) FindAll() ([]*entities.Region, error) {
	var dbRegion []db.Region

	err := r.db.Find(&dbRegion).Error

	if err != nil {
		return nil, err
	}

	Region := make([]*entities.Region, len(dbRegion))

	for i, dbRegion := range dbRegion {
		Region[i] = fromDBRegion(&dbRegion)
	}

	return Region, nil
}
