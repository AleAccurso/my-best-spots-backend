package dtos

import (
	"my-best-spots-backend/enums"
	"time"

	"github.com/google/uuid"
)

type SpotCreateResDTO struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`

	Name         string                `json:"name"`
	CategoryId   uuid.UUID             `json:"category_id"`
	AddressId      uuid.UUID         `json:"address_id"`
	Latitude     float32               `json:"latitude"`
	Longitude    float32               `json:"longitude"`
	MinAuthGroup enums.SpotAccessRight `json:"min_auth_group"`
}