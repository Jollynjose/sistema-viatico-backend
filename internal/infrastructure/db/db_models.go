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
	Provinces      []Province     `json:"provinces" gorm:"foreignKey:RegionID;constraint:OnDelete:CASCADE;REFERENCES:Identifier"`
	Municipalities []Municipality `json:"municipalities" gorm:"foreignKey:RegionID;constraint:OnDelete:CASCADE;REFERENCES:Identifier"`
}

type Province struct {
	Base
	Name           string         `json:"name" gorm:"not null"`
	Identifier     string         `json:"identifier" gorm:"uniqueIndex;not null"`
	RegionID       string         `json:"region_id" gorm:"not null"`
	Municipalities []Municipality `json:"municipalities" gorm:"foreignKey:ProvinceID;constraint:OnDelete:CASCADE;REFERENCES:Identifier"`
}

type Municipality struct {
	Base
	Name       string `json:"name" gorm:"not null"`
	Identifier string `json:"identifier" gorm:"uniqueIndex;not null"`
	ProvinceID string `json:"province_id" gorm:"not null;uniqueIndex:idx_province_municipality"`
	RegionID   string `json:"region_id" gorm:"not null;uniqueIndex:idx_region_municipality"`
}
