package repositories

import (
	"context"
	"my-best-spots-backend/entities"
	"my-best-spots-backend/enums"
)

type ISpotRepository interface {
	InsertSpot(context context.Context, spotEntity entities.SpotEntity) (*entities.SpotEntity, error)
	GetSpots(context context.Context, page *int, size *int) ([]entities.SpotPreloadedEntity, error)
	GetAvailableCountries(context context.Context, roles []enums.SpotAccessRight) ([]entities.CountryEntity, error)
	GetAvailableRegionsByCountryCode(context context.Context, countryCode string, roles []enums.SpotAccessRight) ([]entities.RegionEntity, error)
}
