package controllers

import (
	"errors"
	"my-best-spots-backend/constants"
	"my-best-spots-backend/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	usecases usecases.Usecase
}

func InitialiseCategoryController(usecases usecases.Usecase) CategoryController {
	return CategoryController{usecases: usecases}
}

func (controller CategoryController) GetCategories(c *gin.Context) {

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

	categories, err := controller.usecases.CategoryUsecase.GetCategories(c, pagePtr, sizePtr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, categories)
}

func (controller CategoryController) GetCategoryByKey(c *gin.Context) {

	categoryKey := c.Param("category_key")
	if categoryKey == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constants.MISSING_PARAM+"category_key").Error())
		return
	}

	category, err := controller.usecases.CategoryUsecase.GetCategoryByKey(c, categoryKey)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, category)
}
