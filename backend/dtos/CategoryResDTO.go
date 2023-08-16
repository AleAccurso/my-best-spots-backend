package dtos

import (
	"github.com/google/uuid"
)

type CategoryResDTO struct {
	Id uuid.UUID `bson:"id,omitempty" json:"id"`

	CategoryName string `bson:"category_name,omitempty" json:"category_name"`
	CategoryKey  string `bson:"category_key,omitempty" json:"category_key"`
}
