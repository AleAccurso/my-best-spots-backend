package routers

import (
	"new-rating-movies-go-backend/controllers"
	middlewares "new-rating-movies-go-backend/middelwares"

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
	api.GET("/me", router.authMiddleware.Authorize(router.controller.AuthController.GetMe))
	api.POST("/login", router.controller.AuthController.Login)
	api.GET("/logout", router.authMiddleware.Authorize(router.controller.AuthController.Logout))
	api.POST("/register", router.controller.AuthController.Register)

	// User
	api.GET("/users", router.authMiddleware.Authorize(router.controller.UserController.GetUsers, "admin"))
	api.GET("/users/:userId", router.authMiddleware.Authorize(router.controller.UserController.GetUserById))
	api.POST("/users/:userId", router.authMiddleware.Authorize(router.controller.UserController.UpdateUserById))
	api.DELETE("/users/:userId", router.authMiddleware.Authorize(router.controller.UserController.DeleteUserById))

	// Run the engine
	if err := router.engine.Run(":8010"); err != nil {
		return err
	}

	return nil
}
