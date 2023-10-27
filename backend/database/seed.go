package database

import (
	"my-best-spots-backend/database/models"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	category1 := models.Category{
		Id:           uuid.New(),
		CategoryName: "Cafe - Bar",
		CategoryKey:  "cafe-bar",
	}
	category2 := models.Category{
		Id:           uuid.New(),
		CategoryName: "Gas Station",
		CategoryKey:  "gas-station",
	}
	category3 := models.Category{
		Id:           uuid.New(),
		CategoryName: "Hosting",
		CategoryKey:  "hosting",
	}
	category4 := models.Category{
		Id:           uuid.New(),
		CategoryName: "Leisure",
		CategoryKey:  "leisure",
	}
	category5 := models.Category{
		Id:           uuid.New(),
		CategoryName: "Medical",
		CategoryKey:  "medical",
	}
	category6 := models.Category{
		Id:           uuid.New(),
		CategoryName: "Private Beach",
		CategoryKey:  "private-beach",
	}
	category7 := models.Category{
		Id:           uuid.New(),
		CategoryName: "Producer",
		CategoryKey:  "producer",
	}
	category8 := models.Category{
		Id:           uuid.New(),
		CategoryName: "Public Beach",
		CategoryKey:  "public-beach",
	}
	category9 := models.Category{
		Id:           uuid.New(),
		CategoryName: "Religious Site",
		CategoryKey:  "religious-site",
	}
	category10 := models.Category{
		Id:           uuid.New(),
		CategoryName: "Restaurant",
		CategoryKey:  "restaurant",
	}
	category11 := models.Category{
		Id:           uuid.New(),
		CategoryName: "Shop",
		CategoryKey:  "shop",
	}
	category12 := models.Category{
		Id:           uuid.New(),
		CategoryName: "Thermal",
		CategoryKey:  "thermal",
	}
	category13 := models.Category{
		Id:           uuid.New(),
		CategoryName: "Tourism",
		CategoryKey:  "tourism",
	}

	categories := []*models.Category{&category1, &category2, &category3, &category4, &category5, &category6, &category7, &category8, &category9, &category10, &category11, &category12, &category13}
	db.Model(&models.Category{}).Create(&categories)

	address1 := models.Address{
		Id:           uuid.New(),
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
		DeletedAt:    nil,
		Street:       "Via Centrale",
		StreetNumber: lo.ToPtr("44"),
		PostalCode:   "65020",
		City:         "Musellaro",
		Region:       "Abruzzo",
		CountryName:  "Italy",
		CountryCode:  "IT",
	}
	address2 := models.Address{
		Id:           uuid.New(),
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
		DeletedAt:    nil,
		Street:       "Boulevard Sylvain Dupuis",
		StreetNumber: lo.ToPtr("356"),
		PostalCode:   "1070",
		City:         "Anderlecht",
		Region:       "Bruxelles",
		CountryName:  "Belgium",
		CountryCode:  "BE",
	}
	address3 := models.Address{
		Id:           uuid.New(),
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
		DeletedAt:    nil,
		Street:       "Via Tiburtina Valeria",
		StreetNumber: lo.ToPtr("16"),
		PostalCode:   "65024",
		City:         "Manopello",
		Region:       "Abruzzo",
		CountryName:  "Italy",
		CountryCode:  "IT",
	}

	addresses := []*models.Address{&address1, &address2, &address3}
	db.Model(&models.Address{}).Create(&addresses)

	spot1 := models.Spot{
		Id:           uuid.New(),
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
		DeletedAt:    nil,
		Name:         "Home",
		AddressId:    address1.Id,
		Address:      address1,
		CategoryId:   category3.Id,
		Category:     category3,
		Latitude:     42.19240188598633,
		Longitude:    13.953117370605469,
		MinAuthGroup: "LOGGED_USERS",
	}
	spot2 := models.Spot{
		Id:           uuid.New(),
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
		DeletedAt:    nil,
		Name:         "Anderlecht",
		AddressId:    address2.Id,
		Address:      address2,
		CategoryId:   category3.Id,
		Category:     category3,
		Latitude:     50.83761215209961,
		Longitude:    4.286291122436523,
		MinAuthGroup: "ADMIN",
	}
	spot3 := models.Spot{
		Id:           uuid.New(),
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
		DeletedAt:    nil,
		Name:         "Osteria Del Contadino",
		AddressId:    address3.Id,
		Address:      address3,
		CategoryId:   category10.Id,
		Category:     category10,
		Latitude:     42.31293,
		Longitude:    14.06208,
		MinAuthGroup: "LOGGED_USERS",
	}

	spots := []*models.Spot{&spot1, &spot2, &spot3}
	db.Model(&models.Spot{}).Create(&spots)
}
