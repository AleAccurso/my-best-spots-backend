package mappers

import (
	"my-best-spots-backend/entities"
	"my-best-spots-backend/models"
)

func AddressModelToEntity(model models.Address) entities.AddressEntity {
	return entities.AddressEntity{
		Id:           model.Id,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
		DeletedAt:    model.DeletedAt,
		Street:       model.Street,
		StreetNumber: model.StreetNumber,
		PostalCode:   model.PostalCode,
		City:         model.City,
		Region:       model.Region,
		CountryName:  model.CountryName,
		CountryCode:  model.CountryCode,
	}
}
