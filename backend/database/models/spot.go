package models

import (
	"time"

	"github.com/google/uuid"
)

type Spot struct {
	Id        uuid.UUID  `json:"id" gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;default:now();not null"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;default:now();not null"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`

	Name         string    `json:"name" gorm:"column:name;type:varchar(64);not null"`
	AddressId    uuid.UUID `json:"address_id" gorm:"column:address_id;not null"`
	Address      Address
	CategoryId   uuid.UUID `json:"category_id" gorm:"column:category_id;not null"`
	Category     Category
	Latitude     float32 `json:"latitude" gorm:"column:latitude;min:-90;max:90;not null;index:idx_spot_location,unique"`
	Longitude    float32 `json:"longitude" gorm:"column:longitude;min:-180;max:180;not null;index:idx_spot_location,unique"`
	MinAuthGroup string  `json:"min_auth_group" gorm:"column:min_auth_group;type:varchar(12);not null"`
}
