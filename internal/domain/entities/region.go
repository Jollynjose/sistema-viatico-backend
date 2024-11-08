package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Region struct {
	Id         uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Name       string
	Identifier string
	Code       string
}

func (r *Region) validate() error {
	if r.Name == "" {
		return errors.New("first name is required")
	}

	if r.Identifier == "" {
		return errors.New("identifier is required")
	}

	if r.Code == "" {
		return errors.New("code is required")
	}

	return nil
}

func NewRegion(Name, Identifier, Code string) *Region {
	region := &Region{
		Id:         uuid.New(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Name:       Name,
		Identifier: Identifier,
		Code:       Code,
	}
	return region
}
