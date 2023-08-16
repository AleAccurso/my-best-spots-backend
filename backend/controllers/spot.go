package controllers

import (
	"errors"
	"my-best-spots-backend/constants"
	"my-best-spots-backend/dtos"
	"my-best-spots-backend/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SpotController struct {
	usecases usecases.Usecase
}

func InitialiseSpotController(usecases usecases.Usecase) SpotController {
	return SpotController{usecases: usecases}
}

func (controller SpotController) GetAvailableSpots(c *gin.Context) {
	var pagePtr, sizePtr *int

	page := c.Query("page")
	if page != "" {
		parsed_page, err := strconv.Atoi(page)
		if err != nil {
			pagePtr = nil
		} else {
			pagePtr = &parsed_page
		}
	}

	size := c.Query("size")
	if size != "" {
		parsed_size, err := strconv.Atoi(size)
		if err != nil {
			sizePtr = nil
		} else {
			sizePtr = &parsed_size
		}
	}

	println(c)

	spots, err := controller.usecases.SpotUsecase.GetAvailableSpots(c, pagePtr, sizePtr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, spots)
}

func (controller SpotController) GetSpotById(c *gin.Context) {
	spotId := c.Param("spot_id")
	if spotId == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constants.MISSING_PARAM+"spot_id").Error())
		return
	}

	spotUUID, err := uuid.Parse(spotId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constants.BAD_PARAMS+"spot_id").Error())
		return
	}

	spot, err := controller.usecases.SpotUsecase.GetSpotById(c, spotUUID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, spot)
}

func (controller SpotController) AddSpot(c *gin.Context) {
	var spotReqCreateDTO dtos.SpotReqCreateDTO
	if err := c.ShouldBindJSON(&spotReqCreateDTO); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	spot, err := controller.usecases.SpotUsecase.AddSpot(c, spotReqCreateDTO)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, spot)
}
