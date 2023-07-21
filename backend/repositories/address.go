package repositories

import (
	"context"
	"my-best-spots-backend/database/models"
	"my-best-spots-backend/entities"
	"my-best-spots-backend/repositories/mappers"

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
	var address models.Address

	if err := repository.database.Model(&models.Address{}).Where("id = ?", addressId).First(&address).Error; err != nil {
		return nil, err
	}

	addressEntity := mappers.AddressModelToEntity(address)

	return &addressEntity, nil
}

func (repository AddressRepository) InsertAddress(context context.Context, addressEntity entities.AddressEntity) (*entities.AddressEntity, error) {
	address := mappers.AddressEntityToModel(addressEntity)

	err := repository.database.Model(&models.Address{}).Create(&address).Error
	if err != nil {
		return nil, err
	}

	newAddress := mappers.AddressModelToEntity(address)

	return &newAddress, nil
}

func (repository AddressRepository) CheckAddressAlreadyExists(context context.Context, addressEntity entities.AddressEntity) (*entities.AddressEntity, error) {
	var address models.Address

	err := repository.database.Model(&models.Address{}).Where("street = ?", addressEntity.Street).Where("street_number = ?", addressEntity.StreetNumber).Where("postal_code = ?", addressEntity.PostalCode).Where("city = ?", addressEntity.City).Where("region = ?", addressEntity.Region).Where("country_name = ?", addressEntity.CountryName).First(&address).Error
	if err != nil {
		return nil, err
	}

	dbAddressEntity := mappers.AddressModelToEntity(address)

	return &dbAddressEntity, nil
}
