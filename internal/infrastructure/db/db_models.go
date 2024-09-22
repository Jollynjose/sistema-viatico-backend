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
