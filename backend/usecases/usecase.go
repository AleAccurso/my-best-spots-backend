package usecases

import (
	"my-best-spots-backend/repositories"
	"my-best-spots-backend/services"
)

type Usecase struct {
	UsecaseBase
}

func Initialise(repository repositories.Repository, service services.Service) Usecase {
	return Usecase{
		UsecaseBase: UsecaseBase{
			CategoryUsecase: InitialiseCategoryUsecase(repository),
		},
	}
}
