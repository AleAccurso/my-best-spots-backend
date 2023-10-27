package controllers

import (
	"my-best-spots-backend/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CountryController struct {
	usecases usecases.Usecase
}

func InitialiseCountryController(usecases usecases.Usecase) CountryController {
	return CountryController{usecases: usecases}
}

func (controller CountryController) GetAvailableCountries(c *gin.Context) {

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

	countries, err := controller.usecases.CountryUsecase.GetAvailableCountries(c, pagePtr, sizePtr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, countries)
}
