package repository

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/repositories"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"gorm.io/gorm"
)

type GormMunicipalityRepository struct {
	db *gorm.DB
}

func NewGormMunicipalityRepository(db *gorm.DB) repositories.MunicipalityRepository {
	return &GormMunicipalityRepository{db: db}
}

func (r *GormMunicipalityRepository) Create(user *entities.MunicipalityValidated) (*entities.Municipality, error) {
	dbData := toDBMunicipality(user)

	if err := r.db.Create(&dbData).Error; err != nil {
		return nil, err
	}

	return fromDBMunicipality(dbData), nil
}

func (r *GormMunicipalityRepository) FindById(id string) (*entities.Municipality, error) {
	var dbMunicipality db.Municipality

	err := r.db.First(&dbMunicipality, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return fromDBMunicipality(&dbMunicipality), nil
}

func (r *GormMunicipalityRepository) FindByIdentifier(identifier string) (*entities.Municipality, error) {
	var dbMunicipality db.Municipality

	err := r.db.First(&dbMunicipality, "identifier = ?", identifier).Error

	if err != nil {
		return nil, err
	}

	return fromDBMunicipality(&dbMunicipality), nil
}

func (r *GormMunicipalityRepository) FindAll() ([]*entities.Municipality, error) {
	var dbMunicipality []db.Municipality

	err := r.db.Find(&dbMunicipality).Error

	if err != nil {
		return nil, err
	}

	municipalities := make([]*entities.Municipality, len(dbMunicipality))

	for i, dbMunicipality := range dbMunicipality {
		municipalities[i] = fromDBMunicipality(&dbMunicipality)
	}

	return municipalities, nil
}

func (r *GormMunicipalityRepository) FindByRegionCode(regionCode string) ([]*entities.Municipality, error) {
	var dbMunicipality []db.Municipality

	err := r.db.Find(&dbMunicipality, "region_code = ?", regionCode).Error

	if err != nil {
		return nil, err
	}

	municipalities := make([]*entities.Municipality, len(dbMunicipality))

	for i, dbMunicipality := range dbMunicipality {
		municipalities[i] = fromDBMunicipality(&dbMunicipality)
	}

	return municipalities, nil
}
