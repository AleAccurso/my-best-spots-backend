package controllers

import (
	"errors"
	"my-best-spots-backend/constants"
	"my-best-spots-backend/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (controller CategoryController) GetCategoryById(c *gin.Context) {

	categoryId := c.Param("category_id")
	if categoryId == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constants.MISSING_PARAM+"category_id").Error())
		return
	}

	categoryUUID, err := uuid.Parse(categoryId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constants.BAD_PARAMS+"CategoryId"))
		return
	}

	category, err := controller.usecases.CategoryUsecase.GetCategoryById(c, categoryUUID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, category)
}
