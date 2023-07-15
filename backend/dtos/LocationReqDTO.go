package dtos

type LocationReqCreateDTO struct {
	City         string  `bson:"city, omitempty" json:"city"`
	CountryCode  string  `bson:"country_code, omitempty" json:"country_code"`
	CountryName  string  `bson:"country_name, omitempty" json:"country_name"`
	Latitude     float32 `bson:"latitude, omitempty" json:"latitude"`
	Longitude    float32 `bson:"longitude, omitempty" json:"longitude"`
	PostalCode   string  `bson:"postal_code, omitempty" json:"postal_code"`
	Region       string  `bson:"region, omitempty" json:"region"`
	Street       string  `bson:"street, omitempty" json:"street"`
	StreetNumber *string `bson:"street_number, omitempty" json:"street_number"`
}
