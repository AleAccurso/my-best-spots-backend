package models

import (
	"time"

	"github.com/google/uuid"
)

type Spot struct {
	Id        uuid.UUID  `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:now();not null"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:now();not null"`
	DeletedAt *time.Time `json:"deleted_at"`

	Name         string    `json:"name" gorm:"type:varchar(64);not null"`
	AddressId    uuid.UUID `json:"address_id" gorm:"not null"`
	Address      Address
	CategoryId   uuid.UUID `json:"category_id" gorm:"not null"`
	Category     Category
	Latitude     float32 `json:"latitude" gorm:"min:-90;max:90;not null;index:idx_spot_location,unique"`
	Longitude    float32 `json:"longitude" gorm:"min:-180;max:180;not null;index:idx_spot_location,unique"`
	MinAuthGroup string  `json:"min_auth_group" gorm:"type:varchar(12);not null"`
}
