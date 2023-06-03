package mappers

import (
	"my-best-spots-backend/dtos"
	"my-best-spots-backend/models"
)

func CategoryModelToResDTO(model models.Category) dtos.CategoryResDTO {
	return dtos.CategoryResDTO{
		Id: model.Id,

		CategoryName:   model.CategoryName,
		CategoryKey:    model.CategoryKey,
		SvgIconContent: string(model.SvgIconContent),
	}
}

func CategoryModelsToResDTOs(models []models.Category) []dtos.CategoryResDTO {
	dtos := make([]dtos.CategoryResDTO, len(models))

	for i, model := range models {
		dtos[i] = CategoryModelToResDTO(model)
	}

	return dtos
}
