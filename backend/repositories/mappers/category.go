package mappers

import (
	"my-best-spots-backend/database/models"
	"my-best-spots-backend/entities"
)

func CategoryModelToEntity(model models.Category) entities.CategoryEntity {
	return entities.CategoryEntity{
		Id:           model.Id,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
		DeletedAt:    model.DeletedAt,
		CategoryName: model.Name,
		CategoryKey:  model.Key,
	}
}

func CategoryModelsToEntities(models []models.Category) []entities.CategoryEntity {
	entities := make([]entities.CategoryEntity, len(models))

	for i, model := range models {
		entities[i] = CategoryModelToEntity(model)
	}

	return entities
}
