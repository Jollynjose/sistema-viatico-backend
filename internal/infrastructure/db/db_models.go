package db

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Roles string

const (
	Admin_Role   Roles = "admin"
	General_Role Roles = "general"
	Manager_Role Roles = "manager"
)

type FuelPriceType string

const (
	Diesel_FuelPriceType   FuelPriceType = "diesel"
	Gasoline_FuelPriceType FuelPriceType = "gasoline"
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
	FirstName               string              `json:"first_name" gorm:"size:255;not null"`
	LastName                string              `json:"last_name" gorm:"size:255;not null"`
	Email                   string              `json:"email" gorm:"size:255;not null;unique"`
	Password                string              `json:"password" gorm:"not null"`
	Role                    Roles               `json:"role" gorm:"type:enum('admin', 'manager', 'general');not null"`
	JobPositionId           string              `json:"job_position_id" gorm:"not null"`
	JobPosition             JobPosition         `json:"job_position" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	JobPostionSpecification *string             `json:"job_position_specification" gorm:"size:255"`
	TravelHistories         []UserTravelHistory `json:"travel_histories" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;References:ID"`
}

type JobPosition struct {
	Base
	Name                 string               `json:"name" gorm:"size:255;not null"`
	JobPositionHistories []JobPositionHistory `json:"job_position_histories" gorm:"foreignKey:JobPositionId;constraint:OnDelete:CASCADE;References:ID"`
	Users                []User               `json:"users" gorm:"foreignKey:JobPositionId"`
}

type JobPositionHistory struct {
	Base
	Lunch           float64             `json:"lunch" gorm:"not null"`
	BreakFast       float64             `json:"breakfast" gorm:"not null"`
	Dinner          float64             `json:"dinner" gorm:"not null"`
	JobPositionId   string              `json:"job_position_id" gorm:"not null"`
	TravelHistories []UserTravelHistory `json:"travel_histories" gorm:"foreignKey:JobPositionHistoryId;constraint:OnDelete:CASCADE;References:ID"`
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
	Routes         []Route        `json:"routes" gorm:"foreignKey:StartingPointProvinceId;constraint:OnDelete:CASCADE;References:ID"`
	Stops          []Stop         `json:"stops" gorm:"foreignKey:ProvinceId;constraint:OnDelete:CASCADE;References:ID"`
}

type Municipality struct {
	Base
	Name         string `json:"name" gorm:"not null"`
	Identifier   string `json:"identifier" gorm:"uniqueIndex;not null"`
	ProvinceCode string `json:"province_code" gorm:"not null;index"`
	RegionCode   string `json:"region_code" gorm:"not null;index"`
	Code         string `json:"code" gorm:"not null"`
}

type FuelHistory struct {
	Base
	Price  float64 `json:"price" gorm:"not null"`
	FuelId string  `json:"fuel_id" gorm:"not null"`
}

type Fuel struct {
	Base
	Type    FuelPriceType `json:"type" gorm:"type:enum('diesel', 'gasoline');not null"`
	History []FuelHistory `json:"history" gorm:"foreignKey:FuelId;constraint:OnDelete:CASCADE;References:ID"`
	Name    string        `json:"name" gorm:"not null"`
}

type Route struct {
	Base
	StartingPointProvinceId string   `json:"starting_point_province_id" gorm:"not null"`
	Province                Province `json:"starting_point_province" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Description             string   `json:"description" gorm:"not null"`
	Name                    string   `json:"name" gorm:"not null"`
	TotalKms                int      `json:"total_kms" gorm:"not null"`
	Stops                   []Stop   `json:"stops" gorm:"foreignKey:RouteId;constraint:OnDelete:CASCADE;References:ID"`
}

type Stop struct {
	Base
	Order      int      `json:"order" gorm:"not null"`
	RouteId    string   `json:"route_id" gorm:"not null"`
	Route      Route    `json:"route" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProvinceId string   `json:"province_id" gorm:"not null"`
	Province   Province `json:"province" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Toll struct {
	Base
	Price float64 `json:"price" gorm:"not null"`
	Order int     `json:"order" gorm:"not null"`
}

type UserTravelHistory struct {
	Base
	UserId               string             `json:"user_id" gorm:"not null"`
	User                 User               `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TravelExpense        TravelExpense      `json:"travel_expense" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TravelExpenseId      string             `json:"travel_expense_id" gorm:"not null"`
	JobPositionHistoryId string             `json:"job_position_history_id" gorm:"not null"`
	JobPositionHistory   JobPositionHistory `json:"job_position_history" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TotalPrice           float64            `json:"total_price" gorm:"not null"`
	PlusPercentage       float64            `json:"plus_percentage" gorm:"not null"`
}

type TravelExpense struct {
	Base
	FuelHistoryId     string              `json:"fuel_history_id" gorm:"not null"`
	FuelHistory       Fuel                `json:"fuel_history" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TotalPrice        float64             `json:"total_price" gorm:"not null"`
	DepartureDate     time.Time           `json:"departure_date" gorm:"not null"`
	ArrivalDate       time.Time           `json:"arrival_date" gorm:"not null"`
	SolicitudeDate    time.Time           `json:"solicitude_date" gorm:"not null"`
	RouteId           string              `json:"route_id" gorm:"not null"`
	Route             Route               `json:"route" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserTravelHistory []UserTravelHistory `json:"user_travel_history" gorm:"foreignKey:TravelExpenseId;constraint:OnDelete:CASCADE;References:ID"`
}
