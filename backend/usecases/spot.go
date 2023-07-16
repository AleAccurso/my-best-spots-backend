package usecases

import (
	"my-best-spots-backend/dtos"
	"my-best-spots-backend/repositories"
	"my-best-spots-backend/usecases/mappers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SpotUsecase struct {
	repositories repositories.Repository
}

func InitialiseSpotUsecase(repositories repositories.Repository) SpotUsecase {
	return SpotUsecase{
		repositories: repositories,
	}
}

func (usecase SpotUsecase) GetAvailableSpots(c *gin.Context, page *int, size *int) (*dtos.SpotPagingResDTO, error) {
	return nil, nil
}

func (usecase SpotUsecase) GetSpotById(c *gin.Context, spotId uuid.UUID) (*dtos.SpotResDTO, error) {
	return nil, nil
}

func (usecase SpotUsecase) GetSpotAddress(c *gin.Context, spotUUID uuid.UUID, addressUUID uuid.UUID) (*dtos.SpotResDTO, error) {
	return nil, nil
}

func (usecase SpotUsecase) AddSpot(c *gin.Context, spot dtos.SpotReqCreateDTO) (*dtos.SpotResDTO, error) {
	spotEntity, addressEntity := mappers.SpotReqCreateToSpotAddressEntities(spot)

	// Get Category
	categoryEntity, err := usecase.repositories.CategoryRepository.GetCategoryByKey(c, spot.Category)
	if err != nil {
		return nil, err
	}
	spotEntity.CategoryId = categoryEntity.Id

	// Insert Address if needed
	addressEntityFromDB, err := usecase.repositories.AddressRepository.CheckAddressAlreadyExists(c, addressEntity)
	if err != nil {
		addressEntityFromDB, err = usecase.repositories.AddressRepository.InsertAddress(c, addressEntity)
		if err != nil {
			return nil, err
		}
	}
	spotEntity.AddressId = addressEntityFromDB.Id

	// Call repository
	newSpotEntity, err := usecase.repositories.SpotRepository.InsertSpot(c, spotEntity)
	if err != nil {
		return nil, err
	}

	newSpotDTO := mappers.SpotCreateEntityToDTO(*newSpotEntity)

	return &newSpotDTO, nil
}
