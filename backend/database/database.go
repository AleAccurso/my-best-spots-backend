package database

import (
	"errors"
	"fmt"
	"my-best-spots-backend/constants"
	"my-best-spots-backend/database/models"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initialise() (*gorm.DB, error) {

	if err := godotenv.Load("../.env"); err != nil {
		return nil, errors.New("database: No .env file found")
		// log.Println("No .env file found")
	}

	// Check connection information
	PG_HOST := os.Getenv("POSTGRES_HOST_ADDRESS")
	PG_PORT := os.Getenv("POSTGRES_HOST_PORT")
	PG_USER := os.Getenv("POSTGRES_DOCKER_USER")
	PG_PASSWORD := os.Getenv("POSTGRES_DOCKER_PASSWORD")
	PG_DB_NAME := os.Getenv("POSTGRES_DOCKER_DB_NAME")

	if PG_HOST == "" || PG_PORT == "" || PG_USER == "" || PG_PASSWORD == "" || PG_DB_NAME == "" {
		return nil, errors.New(constants.MISSING_DB_CONNECTION_PARAMS)
	}

	dsn := url.URL{
		User:     url.UserPassword(PG_USER, PG_PASSWORD),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%s", PG_HOST, PG_PORT),
		Path:     PG_DB_NAME,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})
	if err != nil {
		return nil, errors.New(constants.UNABLE_TO_DO_ACTION + "connect-to-database")
	}

	db.AutoMigrate(&models.Address{}, &models.Category{}, &models.Spot{}, &models.User{})

	fmt.Println("Successfully connected!")

	Seed(db)

	return db, nil
}
