package main

import (
	"fmt"
	"my-best-spots-backend/controllers"
	db "my-best-spots-backend/database"
	middlewares "my-best-spots-backend/middelwares"
	"my-best-spots-backend/repositories"
	"my-best-spots-backend/routers"
	"my-best-spots-backend/services"
	"my-best-spots-backend/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	// Initialises the router
	engine := gin.Default()

	// Connects to the database
	database, err := db.Initialise()
	if err != nil {
		return fmt.Errorf("router: %s", err)
	}

	// Creates the repository container
	repository := repositories.Initialise(database)

	// Creates the service container
	service := services.Initialise(repository)

	// Creates the usecase container
	usecase := usecases.Initialise(repository, service)

	// Creates the controller container
	controller := controllers.Initialise(usecase, service)

	// Creates middlewares
	middleware := middlewares.Initialise()

	// Creates the routes container
	router := routers.Initialise(engine, middleware.AuthMiddleware, controller)

	// Start the router
	if err := router.Run(); err != nil {
		return fmt.Errorf("router: %s", err)
	}

	return nil
}
