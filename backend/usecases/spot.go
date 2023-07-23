package usecases

import (
	"errors"
	"math"
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

func (usecase SpotUsecase) GetAvailableSpots(c *gin.Context, page *int, size *int) (*dtos.SpotPreloadedPagingResDTO, error) {

	var nbPages, pageInt, sizeInt int8

	if page == nil {
		pageInt = 1
	}

	spotEntities, err := usecase.repositories.SpotRepository.GetSpots(c, page, size)
	if err != nil {
		return nil, err
	}

	if size != nil {
		nbPages = int8(math.Ceil(float64(len(spotEntities)) / float64(sizeInt)))
	} else {
		nbPages = 1
		if page == nil {
			sizeInt = int8(len(spotEntities))
		}
	}

	if nbPages == 0 {
		nbPages = 1
	}

	if pageInt > nbPages-1 {
		pageInt = nbPages
	}


	spotDTOs := mappers.SpotPreloadedEntitiesToPreloadedResDTOs(spotEntities)

	pagingSpots := dtos.SpotPreloadedPagingResDTO{
		Page:      pageInt,
		Size:      sizeInt,
		NbPages:   nbPages,
		NbResults: int16(len(spotDTOs)),
		Data:      spotDTOs,
	}

	return &pagingSpots, nil


}

func (usecase SpotUsecase) GetSpotById(c *gin.Context, spotId uuid.UUID) (*dtos.SpotResDTO, error) {
	return nil, nil
}

func (usecase SpotUsecase) AddSpot(c *gin.Context, spot dtos.SpotReqCreateDTO) (*dtos.SpotResDTO, error) {
	// Validate location
	if spot.Location.Latitude < -90 || spot.Location.Latitude > 90 || spot.Location.Longitude < -180 || spot.Location.Longitude > 180 {
		return nil, errors.New("invalid Location")
	}

	// Map spot creation request to Entity
	spotEntity, addressEntity := mappers.SpotReqCreateToSpotAddressEntities(spot)

	// Get Category
	categoryEntity, err := usecase.repositories.CategoryRepository.GetCategoryByKey(c, spot.Category)
	if err != nil {
		return nil, err
	}
	spotEntity.CategoryId = categoryEntity.Id

	// Insert Address if needed
	newSpotAddressEntity, err := usecase.repositories.AddressRepository.CheckAddressAlreadyExists(c, addressEntity)
	if err != nil {
		newSpotAddressEntity, err = usecase.repositories.AddressRepository.InsertAddress(c, addressEntity)
		if err != nil {
			return nil, err
		}
	}
	spotEntity.AddressId = newSpotAddressEntity.Id

	// Call repository
	newSpotEntity, err := usecase.repositories.SpotRepository.InsertSpot(c, spotEntity)
	if err != nil {
		return nil, err
	}

	newSpotDTO := mappers.SpotCreateEntityToDTO(*newSpotEntity)
	newSpotDTO.Address = mappers.AddressEntityToResDTO(*newSpotAddressEntity)
	newSpotDTO.Category = mappers.CategoryEntityToResDTO(*categoryEntity)

	return &newSpotDTO, nil
}
