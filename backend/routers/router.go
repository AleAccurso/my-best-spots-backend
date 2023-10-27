package routers

import (
	"my-best-spots-backend/controllers"
	middlewares "my-best-spots-backend/middelwares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine         *gin.Engine
	authMiddleware middlewares.IAuthMiddleware
	controller     controllers.Controller
}

func Initialise(engine *gin.Engine, authMiddleware middlewares.IAuthMiddleware, controller controllers.Controller) Router {
	return Router{
		engine:         engine,
		authMiddleware: authMiddleware,
		controller:     controller,
	}
}

func (router Router) Run() error {

	// Setup CORS
	router.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	// Creates the api-group
	api := router.engine.Group("/api")

	////////////////////////////////////////
	//     Initialises all the routers    //
	////////////////////////////////////////

	// Authentication
	// api.GET("/me", router.authMiddleware.Authorize(router.controller.AuthController.GetMe))
	// api.POST("/login", router.controller.AuthController.Login)
	// api.GET("/logout", router.authMiddleware.Authorize(router.controller.AuthController.Logout))

	// Categories
	api.GET("/categories", router.controller.CategoryController.GetCategories)
	api.GET("/category/:category_key", router.controller.CategoryController.GetCategoryByKey)

	// Spot
	api.GET("/spots", router.controller.SpotController.GetAvailableSpots)
	api.GET("/spots/:spot_id", router.controller.SpotController.GetSpotById)
	api.POST("/spots", router.controller.SpotController.AddSpot)

	// Countries
	api.GET("/countries", router.controller.CountryController.GetAvailableCountries)

	// Run the engine
	if err := router.engine.Run(":8010"); err != nil {
		return err
	}

	return nil
}
