package usecases

import (
	"my-best-spots-backend/dtos"

	"github.com/gin-gonic/gin"
)

type ICategoryUsecase interface {
	GetCategories(ctx *gin.Context, page *int, size *int) (*dtos.CategoryPagingResDTO, error)
	GetCategoryByKey(c *gin.Context, categoryKey string) (*dtos.CategoryResDTO, error)
}
