package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Municipality struct {
	Id           uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Code         string
	Name         string
	Identifier   string
	RegionCode   string
	ProvinceCode string
}

func (e *Municipality) validate() error {
	if e.Name == "" {
		return errors.New("first name is required")
	}

	if e.Identifier == "" {
		return errors.New("identifier is required")
	}

	if e.RegionCode == "" {
		return errors.New("region id is required")
	}

	if e.ProvinceCode == "" {
		return errors.New("province id is required")
	}

	if e.Code == "" {
		return errors.New("code is required")
	}

	return nil
}

func NewMunicipality(Name, Identifier, RegionCode, ProvinceCode, Code string) *Municipality {
	municipality := &Municipality{
		Id:           uuid.New(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Name:         Name,
		Identifier:   Identifier,
		RegionCode:   RegionCode,
		ProvinceCode: ProvinceCode,
		Code:         Code,
	}
	return municipality
}
