package mappers

import (
	"my-best-spots-backend/entities"
	"my-best-spots-backend/models"
)

func CategoryModelToEntity(model models.Category) entities.CategoryEntity {
	return entities.CategoryEntity{
		Id:             model.Id,
		CreatedAt:      model.CreatedAt,
		UpdatedAt:      model.UpdatedAt,
		DeletedAt:      model.DeletedAt,
		CategoryName:   model.CategoryName,
		CategoryKey:    model.CategoryKey,
		SvgIconContent: model.SvgIconContent,
	}
}

func CategoryModelsToEntities(models []models.Category) []entities.CategoryEntity {
	entities := make([]entities.CategoryEntity, len(models))

	for i, model := range models {
		entities[i] = CategoryModelToEntity(model)
	}

	return entities
}
