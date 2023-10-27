package usecases

import (
	"my-best-spots-backend/dtos"

	"github.com/gin-gonic/gin"
)

type ICountryUsecase interface {
	GetAvailableCountries(ctx *gin.Context, page *int, size *int) (*dtos.CategoryPagingResDTO, error)
}
