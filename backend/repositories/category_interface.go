package repositories

import (
	"context"
	"my-best-spots-backend/entities"
)

type ICategoryRepository interface {
	GetCategories(context context.Context, page *int, size *int) ([]entities.CategoryEntity, error)
	GetCategoryByKey(context context.Context, categoryKey string) (*entities.CategoryEntity, error)
	CountCategories(context context.Context) (*int64, error)
}
