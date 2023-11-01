package mappers

import (
	"my-best-spots-backend/dtos"
	"my-best-spots-backend/entities"
)

func AddressEntityToResDTO(entity entities.AddressEntity) dtos.AddressResDTO {
	return dtos.AddressResDTO{
		Id:           entity.Id,
		Street:       entity.Street,
		StreetNumber: entity.StreetNumber,
		PostalCode:   entity.PostalCode,
		City:         entity.City,
		RegionName:   entity.RegionName,
		CountryName:  entity.CountryName,
		CountryCode:  entity.CountryCode,
	}
}
