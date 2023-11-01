package dtos

type CountryResDTO struct {
	Name string `bson:"name,omitempty" json:"name"`
	Code string `bson:"code,omitempty" json:"code"`
}
