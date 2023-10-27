package dtos

import (
	"time"

	"github.com/google/uuid"
)

type UserResDTO struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}
