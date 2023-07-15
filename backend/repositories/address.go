package repositories

import (
	"context"
	"my-best-spots-backend/entities"

	"github.com/google/uuid"
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

func (repository AddressRepository) GetAddressById(context context.Context, addressId uuid.UUID) (*entities.AddressEntity, error) {
	return nil, nil
}

func (repository AddressRepository) InsertAddress(context context.Context, addressEntity entities.AddressEntity) (*entities.AddressEntity, error) {
	return nil, nil
}

func (repository AddressRepository) CheckAddressAlreadyExists(context context.Context, addressEntity entities.AddressEntity) (*entities.AddressEntity, error) {
	return nil, nil
}
