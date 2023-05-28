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
			CategoryRepository: InitialiseCategoryRepository(database),
			AuthRepository:     InitialiseAuthRepository(database),
		},
	}
}
