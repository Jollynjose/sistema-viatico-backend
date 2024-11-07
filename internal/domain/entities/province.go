package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Province struct {
	Id         uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Name       string
	Identifier string
	RegionID   string
}

func (p *Province) validate() error {
	if p.Name == "" {
		return errors.New("first name is required")
	}

	if p.Identifier == "" {
		return errors.New("identifier is required")
	}

	if p.RegionID == "" {
		return errors.New("region id is required")
	}
	return nil
}

func NewProvince(Name, Identifier, RegionID string) *Province {
	province := &Province{
		Id:         uuid.New(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Name:       Name,
		Identifier: Identifier,
		RegionID:   RegionID,
	}
	return province
}
