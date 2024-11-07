package repository

import (
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/entities"
	"github.com/Jollynjose/sistema-viatico-backend/internal/domain/repositories"
	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"gorm.io/gorm"
)

type GormProvinceRepository struct {
	db *gorm.DB
}

func NewGormProvinceRepository(db *gorm.DB) repositories.ProvinceRepository {
	return &GormProvinceRepository{db: db}
}

func (r *GormProvinceRepository) Create(user *entities.ProvinceValidated) (*entities.Province, error) {
	dbData := toDBProvince(user)

	if err := r.db.Create(&dbData).Error; err != nil {
		return nil, err
	}

	return fromDBProvince(dbData), nil
}

func (r *GormProvinceRepository) FindById(id string) (*entities.Province, error) {
	var dbProvince db.Province

	err := r.db.First(&dbProvince, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return fromDBProvince(&dbProvince), nil
}

func (r *GormProvinceRepository) FindByIdentifier(identifier string) (*entities.Province, error) {
	var dbProvince db.Province

	err := r.db.First(&dbProvince, "identifier = ?", identifier).Error

	if err != nil {
		return nil, err
	}

	return fromDBProvince(&dbProvince), nil
}

func (r *GormProvinceRepository) FindAll() ([]*entities.Province, error) {
	var dbProvince []db.Province

	err := r.db.Find(&dbProvince).Error

	if err != nil {
		return nil, err
	}

	provinces := make([]*entities.Province, len(dbProvince))

	for i, dbProvince := range dbProvince {
		provinces[i] = fromDBProvince(&dbProvince)
	}

	return provinces, nil
}
