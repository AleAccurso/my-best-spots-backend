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
		Id:   uuid.New(),
		Name: "Cafe - Bar",
		Key:  "cafe-bar",
	}
	category2 := models.Category{
		Id:   uuid.New(),
		Name: "Gas Station",
		Key:  "gas-station",
	}
	category3 := models.Category{
		Id:   uuid.New(),
		Name: "Hosting",
		Key:  "hosting",
	}
	category4 := models.Category{
		Id:   uuid.New(),
		Name: "Leisure",
		Key:  "leisure",
	}
	category5 := models.Category{
		Id:   uuid.New(),
		Name: "Medical",
		Key:  "medical",
	}
	category6 := models.Category{
		Id:   uuid.New(),
		Name: "Private Beach",
		Key:  "private-beach",
	}
	category7 := models.Category{
		Id:   uuid.New(),
		Name: "Producer",
		Key:  "producer",
	}
	category8 := models.Category{
		Id:   uuid.New(),
		Name: "Public Beach",
		Key:  "public-beach",
	}
	category9 := models.Category{
		Id:   uuid.New(),
		Name: "Religious Site",
		Key:  "religious-site",
	}
	category10 := models.Category{
		Id:   uuid.New(),
		Name: "Restaurant",
		Key:  "restaurant",
	}
	category11 := models.Category{
		Id:   uuid.New(),
		Name: "Shop",
		Key:  "shop",
	}
	category12 := models.Category{
		Id:   uuid.New(),
		Name: "Thermal",
		Key:  "thermal",
	}
	category13 := models.Category{
		Id:   uuid.New(),
		Name: "Tourism",
		Key:  "tourism",
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
		RegionName:   "Abruzzo",
		RegionKey:    "abruzzo",
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
		RegionName:   "Bruxelles",
		RegionKey:    "bruxelles",
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
		RegionName:   "Abruzzo",
		RegionKey:    "abruzzo",
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
