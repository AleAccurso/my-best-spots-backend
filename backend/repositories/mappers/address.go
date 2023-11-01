package mappers

import (
	"my-best-spots-backend/database/models"
	"my-best-spots-backend/entities"
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
		RegionName:   model.RegionName,
		CountryName:  model.CountryName,
		CountryCode:  model.CountryCode,
	}
}

func AddressEntityToModel(entity entities.AddressEntity) models.Address {
	return models.Address{
		Street:       entity.Street,
		StreetNumber: entity.StreetNumber,
		PostalCode:   entity.PostalCode,
		City:         entity.City,
		RegionName:   entity.RegionName,
		CountryName:  entity.CountryName,
		CountryCode:  entity.CountryCode,
	}
}
