package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	Id        uuid.UUID  `json:"id" gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;default:now();not null"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;default:now();not null"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`

	Name string `json:"name" gorm:"column:name;type:varchar(64);not null;index:idx_category,unique"`
	Key  string `json:"key" gorm:"column:key;type:varchar(64);not null;unique;index:idx_category,unique"`
}
