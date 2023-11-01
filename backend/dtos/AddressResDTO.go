package dtos

import (
	"github.com/google/uuid"
)

type AddressResDTO struct {
	Id uuid.UUID `json:"id"`

	Street       string  `json:"street"`
	StreetNumber *string `json:"street_number"`
	PostalCode   string  `json:"postal_code"`
	City         string  `json:"city"`
	RegionName   string  `json:"region_name"`
	CountryName  string  `json:"country_name"`
	CountryCode  string  `json:"country_code"`
}
