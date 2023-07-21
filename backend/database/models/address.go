package models

import (
	"time"

	"github.com/google/uuid"
)

type Address struct {
	Id        uuid.UUID  `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:now();not null"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:now();not null"`
	DeletedAt *time.Time `json:"deleted_at"`

	Street       string  `json:"street" gorm:"type:varchar(128);not null"`
	StreetNumber *string `json:"street_number" gorm:"type:varchar(16);not null"`
	PostalCode   string  `json:"postal_code" gorm:"type:varchar(16);not null"`
	City         string  `json:"city" gorm:"type:varchar(64);not null"`
	Region       string  `json:"region" gorm:"type:varchar(32);not null"`
	CountryName  string  `json:"country_name" gorm:"type:varchar(32);not null"`
	CountryCode  string  `json:"country_code" gorm:"type:varchar(3);not null"`
}
