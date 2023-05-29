package repositories

import (
	"context"
	"my-best-spots-backend/models"

	"github.com/google/uuid"
)

type ICategoryRepository interface {
	GetCategories(context context.Context, page *int, size *int, categoryKeyFilter *string) ([]models.Category, error)
	GetCategoryById(context context.Context, CategoryId uuid.UUID) (*models.Category, error)
	CountCategories(context context.Context, categoryKeyFilter *string) (*int64, error)
}
