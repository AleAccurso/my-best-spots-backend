package repositories

import (
	"context"
	"my-best-spots-backend/models"
)

type ICategoryRepository interface {
	GetCategories(context context.Context, page *int, size *int) ([]models.Category, error)
	GetCategoryByKey(context context.Context, categoryKey string) (*models.Category, error)
	CountCategories(context context.Context) (*int64, error)
}
