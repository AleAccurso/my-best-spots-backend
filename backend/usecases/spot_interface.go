package usecases

import (
	"my-best-spots-backend/dtos"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ISpotUsecase interface {
	GetAvailableSpots(ctx *gin.Context, page *int, size *int) (*dtos.SpotPreloadedPagingResDTO, error)
	GetSpotById(c *gin.Context, spotId uuid.UUID) (*dtos.SpotResDTO, error)
	AddSpot(c *gin.Context, spot dtos.SpotReqCreateDTO) (*dtos.SpotCreateResDTO, error)
	GetAvailableCountries(ctx *gin.Context) (*dtos.CountryListResDTO, error)
	GetAvailableRegions(ctx *gin.Context, countryCode string) (*dtos.RegionListResDTO, error)
}
