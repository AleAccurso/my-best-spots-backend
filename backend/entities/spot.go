package entities

import (
	"my-best-spots-backend/enums"
	"time"

	"github.com/google/uuid"
)

type SpotEntity struct {
	Id        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty" json:"deleted_at"`

	Name         string                `json:"name"`
	CategoryId   uuid.UUID             `json:"category_id"`
	AddressId    uuid.UUID             `json:"address_id"`
	Latitude     float32               `json:"latitude"`
	Longitude    float32               `json:"longitude"`
	MinAuthGroup enums.SpotAccessRight `json:"min_auth_group"`
	ArchivedAt   *time.Time            `json:"archived_at"`
}

type SpotPreloadedEntity struct {
	Id        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	Name         string                `json:"name"`
	Category   CategoryEntity             `json:"category"`
	Address    AddressEntity             `json:"address"`
	Latitude     float32               `json:"latitude"`
	Longitude    float32               `json:"longitude"`
	MinAuthGroup enums.SpotAccessRight `json:"min_auth_group"`
	ArchivedAt   *time.Time            `json:"archived_at"`
}
