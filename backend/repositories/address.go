package repositories

import (
	"gorm.io/gorm"
)

type AddressRepository struct {
	database *gorm.DB
}

func InitialiseAddressRepository(db *gorm.DB) AddressRepository {
	return AddressRepository{
		database: db,
	}
}