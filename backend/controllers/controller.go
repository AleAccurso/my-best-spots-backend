package controllers

import (
	"my-best-spots-backend/services"
	"my-best-spots-backend/usecases"
)

type Controller struct {
	CategoryController CategoryController
	AuthController     AuthController
	SpotController SpotController
}

func Initialise(usecases usecases.Usecase, services services.Service) Controller {
	return Controller{
		CategoryController: InitialiseCategoryController(usecases),
		AuthController:     InitialiseAuthController(usecases, services),
		SpotController: InitialiseSpotController(usecases),
	}
}
