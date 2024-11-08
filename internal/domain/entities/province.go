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
	Code       string
	Name       string
	Identifier string
	RegionCode string
}

func (p *Province) validate() error {
	if p.Name == "" {
		return errors.New("first name is required")
	}

	if p.Identifier == "" {
		return errors.New("identifier is required")
	}

	if p.RegionCode == "" {
		return errors.New("region id is required")
	}

	if p.Code == "" {
		return errors.New("code is required")
	}

	return nil
}

func NewProvince(Name, Identifier, RegionCode, Code string) *Province {
	province := &Province{
		Id:         uuid.New(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Name:       Name,
		Identifier: Identifier,
		RegionCode: RegionCode,
		Code:       Code,
	}
	return province
}
