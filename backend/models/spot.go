package models

import (
	"time"

	"github.com/google/uuid"
)

type Spot struct {
	Id        uuid.UUID  `bson:"id,omitempty" json:"id"`
	CreatedAt time.Time  `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at,omitempty" json:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty" json:"deleted_at"`

	Name        string     `bson:"name,omitempty" json:"name"`
	CategoryId   uuid.UUID  `bson:"category_id,omitempty" json:"category_id"`
	AddressId    uuid.UUID  `bson:"address_id,omitempty" json:"address_id"`
	Latitude     float32    `bson:"latitude,omitempty" json:"latitude"`
	Longitude    float32    `bson:"longitude,omitempty" json:"longitude"`
	MinAuthGroup string     `bson:"min_auth_group,omitempty" json:"min_auth_group"`
}
