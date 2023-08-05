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
			IconURL:      "",
		},
		{
			CategoryName: "Gas Station",
			CategoryKey:  "gas-station",
			IconURL:      "",
		},
		{
			CategoryName: "Hosting",
			CategoryKey:  "hosting",
			IconURL:      "",
		},
		{
			CategoryName: "Leisure",
			CategoryKey:  "leisure",
			IconURL:      "",
		},
		{
			CategoryName: "Medical",
			CategoryKey:  "medical",
			IconURL:      "",
		},
		{
			CategoryName: "Private Beach",
			CategoryKey:  "private-beach",
			IconURL:      "",
		},
		{
			CategoryName: "Producer",
			CategoryKey:  "producer",
			IconURL:      "",
		},
		{
			CategoryName: "Public Beach",
			CategoryKey:  "public-beach",
			IconURL:      "",
		},
		{
			CategoryName: "Religious Site",
			CategoryKey:  "religious-site",
			IconURL:      "",
		},
		{
			CategoryName: "Restaurant",
			CategoryKey:  "restaurant",
			IconURL:      "",
		},
		{
			CategoryName: "Shop",
			CategoryKey:  "shop",
			IconURL:      "",
		},
		{
			CategoryName: "Thermal",
			CategoryKey:  "thermal",
			IconURL:      "",
		},
		{
			CategoryName: "Tourism",
			CategoryKey:  "tourism",
			IconURL:      "",
		},
	}

			db.Model(&models.Category{}).Create(&categories)
}
