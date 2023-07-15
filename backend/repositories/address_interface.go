package repositories

import (
	"context"
	"my-best-spots-backend/entities"

	"github.com/google/uuid"
)

type IAddressRepository interface {
	InsertAddress(context context.Context, addressEntity entities.AddressEntity) (*entities.AddressEntity, error)
	CheckAddressAlreadyExists(context context.Context, addressEntity entities.AddressEntity) (*entities.AddressEntity, error)
	GetAddressById(context context.Context, addressId uuid.UUID) (*entities.AddressEntity, error)
}
