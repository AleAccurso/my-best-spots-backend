package usecases

import (
	"my-best-spots-backend/dtos"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ICategoryUsecase interface {
	GetCategories(ctx *gin.Context, page *int, size *int, categoryKeyFilter *string) (*dtos.CategoryPagingResDTO, error)
	GetCategoryById(c *gin.Context, categoryId uuid.UUID) (*dtos.CategoryResDTO, error)
	// GetCategoryByEmail(c *gin.Context, email string) (*dtos.CategoryResDTO, error)
}
