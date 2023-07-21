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

func (repository SpotRepository) GetSpots(context context.Context, page *int, size *int) ([]entities.SpotPreloadedEntity, error) {
	var spots []models.Spot

	query := repository.database

	// Paging
	if page != nil && size != nil {
		query = query.Limit(*size).Offset(*page * *size)
	}

	// Run query
	err := query.Preload("Address").Preload("Category").Find(&spots).Error
	if err != nil {
		return nil, err
	}

	spotEntities, err := mappers.SpotModelsToPreloadedEntities(spots)
	if err != nil {
		return nil, err
	}

	return spotEntities, nil
}
