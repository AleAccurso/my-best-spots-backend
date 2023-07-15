package repositories

import (
	"context"
	"my-best-spots-backend/entities"

	"gorm.io/gorm"
)

type SpotRepository struct {
	database *gorm.DB
}

func InitialiseSpotRepository(db *gorm.DB) SpotRepository {
	return SpotRepository{
		database: db,
	}
}

func (repository SpotRepository) InsertSpot(context context.Context, spotEntity entities.SpotEntity) (*entities.SpotEntity, error) {
	return nil, nil
}
