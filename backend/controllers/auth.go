package controllers

import (
	"my-best-spots-backend/services"
	"my-best-spots-backend/usecases"
)

type AuthController struct {
	usecases usecases.Usecase
	services services.Service
}

func InitialiseAuthController(usecases usecases.Usecase, services services.Service) AuthController {
	return AuthController{usecases: usecases, services: services}
}

// func (controller AuthController) Register(c *gin.Context) {

// 	var CategoryReqCreateDTO dtos.CategoryReqCreateDTO

// 	if err := c.ShouldBindBodyWith(&CategoryReqCreateDTO, binding.JSON); err != nil {
// 		c.IndentedJSON(http.StatusBadRequest, errors.New(constants.UNABLE_TO_BIND_BODY).Error())
// 		return
// 	}

// 	newId, err := controller.services.AuthService.Register(c, CategoryReqCreateDTO)
// 	if err != nil {
// 		c.IndentedJSON(http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, newId)
// }

// func (controller AuthController) Login(c *gin.Context) {

// 	var loginReqDTO dtos.LoginReqDTO
// 	if err := c.ShouldBindJSON(&loginReqDTO); err != nil {
// 		c.IndentedJSON(http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	token, err := controller.services.AuthService.Login(c, loginReqDTO)
// 	if err != nil {
// 		c.IndentedJSON(http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, gin.H{"token": token})
// }

// func (controller AuthController) Logout(c *gin.Context) {
// 	controller.services.AuthService.Logout(c)
// 	c.IndentedJSON(http.StatusOK, constants.SUCCESS_ACTION+"logout")
// }

// func (controller AuthController) GetMe(c *gin.Context) {

// 	CategoryEmail, ok := c.MustGet("Category_email").(string)
// 	if !ok {
// 		c.IndentedJSON(http.StatusBadRequest, errors.New("email not in context"))
// 		return
// 	}

// 	Category, err := controller.usecases.CategoryUsecase.GetCategoryByEmail(c, CategoryEmail)
// 	if err != nil {
// 		c.IndentedJSON(http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, Category)

// }
