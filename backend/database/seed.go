package database

import (
	"my-best-spots-backend/database/models"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	categories := []models.Category{
		{
			CategoryName: "Cafe - Bar",
			CategoryKey:  "cafe-bar",
		},
		{
			CategoryName: "Gas Station",
			CategoryKey:  "gas-station",
		},
		{
			CategoryName: "Hosting",
			CategoryKey:  "hosting",
		},
		{
			CategoryName: "Leisure",
			CategoryKey:  "leisure",
		},
		{
			CategoryName: "Medical",
			CategoryKey:  "medical",
		},
		{
			CategoryName: "Private Beach",
			CategoryKey:  "private-beach",
		},
		{
			CategoryName: "Producer",
			CategoryKey:  "producer",
		},
		{
			CategoryName: "Public Beach",
			CategoryKey:  "public-beach",
		},
		{
			CategoryName: "Religious Site",
			CategoryKey:  "religious-site",
		},
		{
			CategoryName: "Restaurant",
			CategoryKey:  "restaurant",
		},
		{
			CategoryName: "Shop",
			CategoryKey:  "shop",
		},
		{
			CategoryName: "Thermal",
			CategoryKey:  "thermal",
		},
		{
			CategoryName: "Tourism",
			CategoryKey:  "tourism",
		},
	}

	db.Model(&models.Category{}).Create(&categories)
}
