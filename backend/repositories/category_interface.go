package repositories

import (
	"context"
	"my-best-spots-backend/models"

	"github.com/google/uuid"
)

type ICategoryRepository interface {
	GetCategories(context context.Context, page *int, size *int) ([]models.Category, error)
	GetCategoryById(context context.Context, CategoryId uuid.UUID) (*models.Category, error)
	// GetCategoryByEmail(context context.Context, email string) (*models.Category, error)
	CountCategories(context context.Context) (*int64, error)
}
