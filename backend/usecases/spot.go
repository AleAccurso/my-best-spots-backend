package usecases

import (
	"my-best-spots-backend/dtos"
	"my-best-spots-backend/repositories"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SpotUsecase struct {
	repositories repositories.Repository
}

func InitialiseSpotUsecase(repositories repositories.Repository) SpotUsecase {
	return SpotUsecase{
		repositories: repositories,
	}
}

func (usecase SpotUsecase) GetAvailableSpots(ctx *gin.Context, page *int, size *int) (*dtos.CategoryPagingResDTO, error){
	return nil, nil
}

func (usecase SpotUsecase) GetSpotById(c *gin.Context, spotId uuid.UUID) (*dtos.CategoryResDTO, error){
	return nil, nil
}

func (usecase SpotUsecase) GetSpotAddress(c *gin.Context, spotUUID uuid.UUID, addressUUID uuid.UUID) (*dtos.CategoryResDTO, error){
	return nil, nil
}

func (usecase SpotUsecase) AddSpot(c *gin.Context, spot dtos.SpotReqCreateDTO) (*dtos.CategoryResDTO, error){
	return nil, nil
}

