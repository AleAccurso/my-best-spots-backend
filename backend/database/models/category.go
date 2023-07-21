package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	Id        uuid.UUID  `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:now();not null"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:now();not null"`
	DeletedAt *time.Time `json:"deleted_at"`

	CategoryName string `json:"category_name" gorm:"type:varchar(64);not null"`
	CategoryKey  string `json:"category_key" gorm:"type:varchar(64);not null"`
	IconURL      string `json:"icon_url" gorm:"type:varchar(256)"`
}
