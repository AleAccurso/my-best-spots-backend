package usecases

import (
	"errors"
	"my-best-spots-backend/dtos"
	"my-best-spots-backend/repositories"

	"github.com/gin-gonic/gin"
)

type CountryUsecase struct {
	repositories repositories.Repository
}

func InitialiseCountryUsecase(repositories repositories.Repository) CountryUsecase {
	return CountryUsecase{
		repositories: repositories,
	}
}

func (usecase CountryUsecase) GetAvailableCountries(ctx *gin.Context, page *int, size *int) (*dtos.CategoryPagingResDTO, error) {
	return nil, errors.New("Method not yet implemented")
}
