package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Municipality struct {
	Id         uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Name       string
	Identifier string
	RegionID   string
	ProvinceID string
}

func (e *Municipality) validate() error {
	if e.Name == "" {
		return errors.New("first name is required")
	}

	if e.Identifier == "" {
		return errors.New("identifier is required")
	}

	if e.RegionID == "" {
		return errors.New("region id is required")
	}

	if e.ProvinceID == "" {
		return errors.New("province id is required")
	}

	return nil
}

func NewMunicipality(Name, Identifier, RegionID, ProvinceID string) *Municipality {
	municipality := &Municipality{
		Id:         uuid.New(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Name:       Name,
		Identifier: Identifier,
		RegionID:   RegionID,
		ProvinceID: ProvinceID,
	}
	return municipality
}
