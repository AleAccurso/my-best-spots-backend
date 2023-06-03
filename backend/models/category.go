package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	Id        uuid.UUID `bson:"id,omitempty" json:"id"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at"`
	DeletedAt time.Time `bson:"deleted_at,omitempty" json:"deleted_at"`

	CategoryName   string `bson:"category_name,omitempty" json:"category_name"`
	CategoryKey    string `bson:"category_key,omitempty" json:"category_key"`
	SvgIconContent []byte `bson:"svg_icon_content,omitempty" json:"svg_icon_content"`
}
