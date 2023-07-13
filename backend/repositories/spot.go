package repositories

import (
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