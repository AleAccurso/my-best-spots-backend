package dtos

import (
	"my-best-spots-backend/enums"
	"time"

	"github.com/google/uuid"
)

type SpotPreloadedResDTO struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`

	Name         string                `json:"name"`
	Category   CategoryResDTO             `json:"category"`
	Address      AddressResDTO         `json:"address"`
	Latitude     float32               `json:"latitude"`
	Longitude    float32               `json:"longitude"`
	MinAuthGroup enums.SpotAccessRight `json:"min_auth_group"`
}
