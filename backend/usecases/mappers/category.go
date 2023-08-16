package mappers

import (
	"my-best-spots-backend/dtos"
	"my-best-spots-backend/entities"
)

func CategoryEntityToResDTO(entity entities.CategoryEntity) dtos.CategoryResDTO {
	return dtos.CategoryResDTO{
		Id: entity.Id,

		CategoryName: entity.CategoryName,
		CategoryKey:  entity.CategoryKey,
	}
}

func CategoryEntitiesToResDTOs(entities []entities.CategoryEntity) []dtos.CategoryResDTO {
	dtos := make([]dtos.CategoryResDTO, len(entities))

	for i, entity := range entities {
		dtos[i] = CategoryEntityToResDTO(entity)
	}

	return dtos
}
