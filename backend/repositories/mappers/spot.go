package mappers

import (
	"errors"
	"my-best-spots-backend/database/models"
	"my-best-spots-backend/entities"
	"my-best-spots-backend/enums"
)

func SpotModelToEntity(model models.Spot) (*entities.SpotEntity, error) {
	minAuthGroup, ok := enums.ParseToSpotAccessRight(model.MinAuthGroup)
	if !ok {
		return nil, errors.New("")
	}
	return &entities.SpotEntity{
		Id:           model.Id,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
		DeletedAt:    model.DeletedAt,
		Name:         model.Name,
		CategoryId:   model.CategoryId,
		AddressId:    model.AddressId,
		Latitude:     model.Latitude,
		Longitude:    model.Longitude,
		MinAuthGroup: minAuthGroup,
	}, nil
}

func SpotEntityToModel(entity entities.SpotEntity) models.Spot {
	return models.Spot{
		Name:         entity.Name,
		CategoryId:   entity.CategoryId,
		AddressId:    entity.AddressId,
		Latitude:     entity.Latitude,
		Longitude:    entity.Longitude,
		MinAuthGroup: string(entity.MinAuthGroup),
	}
}

func SpotModelToPreloadedEntity(model models.Spot) (*entities.SpotPreloadedEntity, error) {
	minAuthGroup, ok := enums.ParseToSpotAccessRight(model.MinAuthGroup)
	if !ok {
		return nil, errors.New("")
	}
	return &entities.SpotPreloadedEntity{
		Id:           model.Id,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
		DeletedAt:    model.DeletedAt,
		Name:         model.Name,
		Category:     CategoryModelToEntity(model.Category),
		Address:      AddressModelToEntity(model.Address),
		Latitude:     model.Latitude,
		Longitude:    model.Longitude,
		MinAuthGroup: minAuthGroup,
	}, nil
}

func SpotModelsToPreloadedEntities(models []models.Spot) ([]entities.SpotPreloadedEntity, error) {
	entities := make([]entities.SpotPreloadedEntity, len(models))

	for i, model := range models {
		entity, err := SpotModelToPreloadedEntity(model)
		if err != nil {
			return nil, err
		}
		entities[i] = *entity
	}

	return entities, nil
}
