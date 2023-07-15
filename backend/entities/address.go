package entities

import (
	"time"

	"github.com/google/uuid"
)

type AddressEntity struct {
	Id        uuid.UUID  `bson:"id,omitempty" json:"id"`
	CreatedAt time.Time  `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at,omitempty" json:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty" json:"deleted_at"`

	Street       string  `bson:"street,omitempty" json:"street"`
	StreetNumber *string `bson:"street_number,omitempty" json:"street_number"`
	PostalCode   string  `bson:"postal_code,omitempty" json:"postal_code"`
	City         string  `bson:"city, omitempty" json:"city"`
	Region       string  `bson:"region, omitempty" json:"region"`
	CountryName  string  `bson:"country_name, omitempty" json:"country_name"`
	CountryCode  string  `bson:"country_code, omitempty" json:"country_code"`
}
