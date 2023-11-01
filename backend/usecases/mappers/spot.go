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

func getMinimumRoleToAccessSpot(spot dtos.SpotReqCreateDTO) enums.SpotAccessRight {
	if spot.Admin {
		return enums.ADMIN
	}
	if spot.LoggedUsers {
		return enums.LOGGED_USERS
	}
	return enums.EVERYONE
}

func SpotCreateEntityToDTO(entity entities.SpotEntity) dtos.SpotCreateResDTO {
	return dtos.SpotCreateResDTO{
		Id:           entity.Id,
		CreatedAt:    entity.CreatedAt,
		Name:         entity.Name,
		CategoryId:   entity.CategoryId,
		AddressId:    entity.AddressId,
		Latitude:     entity.Latitude,
		Longitude:    entity.Longitude,
		MinAuthGroup: entity.MinAuthGroup,
	}
}

func SpotPreloadedEntityToPreloadedResDTO(entity entities.SpotPreloadedEntity) dtos.SpotPreloadedResDTO {
	return dtos.SpotPreloadedResDTO{
		Id:           entity.Id,
		CreatedAt:    entity.CreatedAt,
		Name:         entity.Name,
		Category:     CategoryEntityToResDTO(entity.Category),
		Address:      AddressEntityToResDTO(entity.Address),
		Latitude:     entity.Latitude,
		Longitude:    entity.Longitude,
		MinAuthGroup: entity.MinAuthGroup,
	}
}

func SpotPreloadedEntitiesToPreloadedResDTOs(entities []entities.SpotPreloadedEntity) []dtos.SpotPreloadedResDTO {
	dtos := make([]dtos.SpotPreloadedResDTO, len(entities))

	for i, entity := range entities {
		dtos[i] = SpotPreloadedEntityToPreloadedResDTO(entity)
	}

	return dtos
}

func CountryEntityToResDTO(entity entities.CountryEntity) dtos.CountryResDTO {
	return dtos.CountryResDTO{
		Name: entity.Name,
		Code: entity.Code,
	}
}

func CountryEntitiesToResDTOs(entities []entities.CountryEntity) []dtos.CountryResDTO {
	dtos := make([]dtos.CountryResDTO, len(entities))

	for i, entity := range entities {
		dtos[i] = CountryEntityToResDTO(entity)
	}

	return dtos
}
