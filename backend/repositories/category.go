package repositories

import (
	"context"
	"errors"
	"my-best-spots-backend/constants"
	"my-best-spots-backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	database *gorm.DB
}

func InitialiseCategoryRepository(db *gorm.DB) CategoryRepository {
	return CategoryRepository{
		database: db,
	}
}

func (repository CategoryRepository) GetCategories(context context.Context, page *int, size *int, categoryKeyFilter *string) ([]models.Category, error) {
	var categories []models.Category

	query := repository.database

	// Paging
	if page != nil && size != nil {
		query = query.Limit(*size).Offset(*page * *size)
	}

	// Category Key Filter
	if categoryKeyFilter != nil {
		query = query.Where("category_key = ?", *categoryKeyFilter)
	}

	// Run query
	err := query.Find(&categories).Error
	if err != nil {
		return nil, err
	}



	return categories, nil
}

func (repository CategoryRepository) GetCategoryById(context context.Context, categoryId uuid.UUID) (*models.Category, error) {
	var category models.Category

	query := repository.database

	err := query.Where("id = ?", categoryId.String()).Take(&category).Error
	if err != nil {
		return nil, errors.New(constants.SERVER_ERROR)
	}

	return &category, nil
}

func (repository CategoryRepository) CountCategories(context context.Context, categoryKeyFilter *string) (*int64, error) {
	var count int64

	query := repository.database

	// Category Key Filter
	if categoryKeyFilter != nil {
		query = query.Where("category_key = ?", *categoryKeyFilter)
	}

	// Run query
	err := query.Model(&models.Category{}).Count(&count).Error
	if err != nil {
		return nil, errors.New(constants.SERVER_ERROR)
	}

	return &count, nil
}
