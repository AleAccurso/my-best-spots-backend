package models

import (
	"time"

	"github.com/google/uuid"
)

type Address struct {
	Id        uuid.UUID  `json:"id" gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;default:now();not null"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;default:now();not null"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`

	Street       string  `json:"street" gorm:"column:street;type:varchar(128);not null;index:idx_address,unique"`
	StreetNumber *string `json:"street_number" gorm:"column:street_number;type:varchar(16);not null;index:idx_address,unique"`
	PostalCode   string  `json:"postal_code" gorm:"column:postal_code;type:varchar(16);not null;index:idx_address,unique"`
	City         string  `json:"city" gorm:"column:city;type:varchar(64);not null;index:idx_address,unique"`
	Region       string  `json:"region" gorm:"column:region;type:varchar(32);not null;index:idx_address,unique"`
	CountryName  string  `json:"country_name" gorm:"column:country_name;type:varchar(32);not null;index:idx_address,unique"`
	CountryCode  string  `json:"country_code" gorm:"column:country_code;type:varchar(3);not null;index:idx_address,unique"`
}
