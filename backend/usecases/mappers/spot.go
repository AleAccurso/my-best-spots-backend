package mappers

import (
	"my-best-spots-backend/dtos"
	"my-best-spots-backend/entities"
	"my-best-spots-backend/enums"
)

func SpotReqCreateToSpotAddressEntities(spot dtos.SpotReqCreateDTO) (entities.SpotEntity, entities.AddressEntity) {
	addressEntity := entities.AddressEntity{
		Street:       spot.Location.Street,
		StreetNumber: spot.Location.StreetNumber,
		PostalCode:   spot.Location.PostalCode,
		City:         spot.Location.City,
		Region:       spot.Location.Region,
		CountryName:  spot.Location.CountryName,
		CountryCode:  spot.Location.CountryCode,
	}

	spotEntity := entities.SpotEntity{
		Name:         spot.Name,
		Latitude:     spot.Location.Latitude,
		Longitude:    spot.Location.Longitude,
		MinAuthGroup: getMinimumRoleToAccessSpot(spot),
	}

	return spotEntity, addressEntity
}

func SpotCreateEntityToDTO(entity entities.SpotEntity) dtos.SpotResDTO {
	return dtos.SpotResDTO{
		Id:           entity.Id,
		CreatedAt:    entity.CreatedAt,
		Name:         entity.Name,
		CategoryId:   entity.CategoryId,
		Latitude:     entity.Latitude,
		Longitude:    entity.Longitude,
		MinAuthGroup: entity.MinAuthGroup,
	}
}

func getMinimumRoleToAccessSpot(spot dtos.SpotReqCreateDTO) enums.SpotAccessRight {
	if spot.Admin {
		return enums.ADMIN
	}
	if spot.LoggedUsers {
		return enums.LOGGED_USERS
	}
	return enums.EVERYONE
}
