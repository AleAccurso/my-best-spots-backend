package repositories

import (
	"context"
	"my-best-spots-backend/database/models"
	"my-best-spots-backend/entities"
	"my-best-spots-backend/repositories/mappers"

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
	spot := mappers.SpotEntityToModel(spotEntity)

	err := repository.database.Model(&models.Spot{}).Create(&spot).Error
	if err != nil {
		return nil, err
	}

	newSpot, err := mappers.SpotModelToEntity(spot)
	if err != nil {
		return nil, err
	}

	return newSpot, nil
}
