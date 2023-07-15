package repositories

import (
	"gorm.io/gorm"
)

type Repository struct {
	RepositoryBase
}

func Initialise(database *gorm.DB) Repository {
	return Repository{
		RepositoryBase: RepositoryBase{
			AddressRepository:  InitialiseAddressRepository(database),
			AuthRepository:     InitialiseAuthRepository(database),
			CategoryRepository: InitialiseCategoryRepository(database),
			SpotRepository:     InitialiseSpotRepository(database),
		},
	}
}
