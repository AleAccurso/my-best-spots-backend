package repositories

import (
	"context"
	"my-best-spots-backend/entities"
)

type ISpotRepository interface {
	InsertSpot(context context.Context, spotEntity entities.SpotEntity) (*entities.SpotEntity, error)
	GetSpots(context context.Context, page *int, size *int) ([]entities.SpotPreloadedEntity, error)
}
