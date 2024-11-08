package db

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base contains common columns for all tables.
type Base struct {
	ID        string    `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}

type User struct {
	Base
	FirstName string `json:"first_name" gorm:"size:255;not null"`
	LastName  string `json:"last_name" gorm:"size:255;not null"`
	Email     string `json:"email" gorm:"size:255;not null;unique"`
	Password  string `json:"password" gorm:"not null"`
}

type Region struct {
	Base
	Name           string         `json:"name" gorm:"not null"`
	Identifier     string         `json:"identifier" gorm:"uniqueIndex;not null"`
	Provinces      []Province     `json:"provinces" gorm:"foreignKey:RegionCode;constraint:OnDelete:CASCADE;REFERENCES:Code"`
	Municipalities []Municipality `json:"municipalities" gorm:"foreignKey:RegionCode;constraint:OnDelete:CASCADE;REFERENCES:Code"`
	Code           string         `json:"code" gorm:"uniqueIndex;not null"`
}

type Province struct {
	Base
	Name           string         `json:"name" gorm:"not null"`
	Identifier     string         `json:"identifier" gorm:"uniqueIndex;not null"`
	RegionCode     string         `json:"region_code" gorm:"not null"`
	Municipalities []Municipality `json:"municipalities" gorm:"foreignKey:ProvinceCode;constraint:OnDelete:CASCADE;REFERENCES:Code"`
	Code           string         `json:"code" gorm:"uniqueIndex;not null"`
}

type Municipality struct {
	Base
	Name         string `json:"name" gorm:"not null"`
	Identifier   string `json:"identifier" gorm:"uniqueIndex;not null"`
	ProvinceCode string `json:"province_code" gorm:"not null;index"`
	RegionCode   string `json:"region_code" gorm:"not null;index"`
	Code         string `json:"code" gorm:"not null"`
}
