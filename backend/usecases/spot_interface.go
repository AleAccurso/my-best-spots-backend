package usecases

import (
	"my-best-spots-backend/dtos"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ISpotUsecase interface {
	GetAvailableSpots(ctx *gin.Context, page *int, size *int) (*dtos.CategoryPagingResDTO, error)
	GetSpotById(c *gin.Context, spotId uuid.UUID) (*dtos.CategoryResDTO, error)
	GetSpotAddress(c *gin.Context, spotUUID uuid.UUID, addressUUID uuid.UUID) (*dtos.CategoryResDTO, error)
	AddSpot(c *gin.Context, spot dtos.SpotReqCreateDTO) (*dtos.CategoryResDTO, error)
}