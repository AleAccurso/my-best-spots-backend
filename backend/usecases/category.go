package usecases

import (
	"math"
	"my-best-spots-backend/dtos"
	"my-best-spots-backend/repositories"
	"my-best-spots-backend/usecases/mappers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CategoryUsecase struct {
	repositories repositories.Repository
}

func InitialiseCategoryUsecase(repositories repositories.Repository) CategoryUsecase {
	return CategoryUsecase{
		repositories: repositories,
	}
}

func (usecase CategoryUsecase) GetCategories(ctx *gin.Context, page *int, size *int, categoryKeyFilter *string) (*dtos.CategoryPagingResDTO, error) {
	
	var nbPages, pageInt, sizeInt int8
	
	if page == nil {
		pageInt = 1
	}
	
	categoriesCount, err := usecase.repositories.CategoryRepository.CountCategories(ctx, categoryKeyFilter)
	if err != nil {
		return nil, err
	}

	if size == nil {
		sizeInt = int8(*categoriesCount)
	}

	if size != nil {
		nbPages = int8(math.Ceil(float64(*categoriesCount) / float64(sizeInt)))
	} else {
		nbPages = 1
	}

	if nbPages == 0 {
		nbPages = 1
	}

	if pageInt > nbPages-1 {
		pageInt = nbPages
	}

	users, err := usecase.repositories.CategoryRepository.GetCategories(ctx, page, size, categoryKeyFilter)
	if err != nil {
		return nil, err
	}

	pagingCategories := dtos.CategoryPagingResDTO{
		Page:      pageInt,
		Size:      sizeInt,
		NbPages:   nbPages,
		NbResults: int16(*categoriesCount),
		Data: mappers.CategoryModelsToResDTOs(users),
	}


	return &pagingCategories, nil
}

func (usecase CategoryUsecase) GetCategoryById(ctx *gin.Context, categoryId uuid.UUID) (*dtos.CategoryResDTO, error) {
	category, err := usecase.repositories.CategoryRepository.GetCategoryById(ctx, categoryId)
	if err != nil {
		return nil, err
	}

	categoryDTO := mappers.CategoryModelToResDTO(*category)

	return &categoryDTO, nil
}