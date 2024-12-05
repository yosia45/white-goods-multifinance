package seeders

import (
	"white-goods-multifinace/models"

	"gorm.io/gorm"
)

func SeedItems(db *gorm.DB) {
	items := []models.Item{
		{
			Name:        "Item 1",
			NormalPrice: 10000000,
		},
		{
			Name:        "Item 2",
			NormalPrice: 20000000,
		},
		{
			Name:        "Item 3",
			NormalPrice: 50000000,
		},
	}

	for _, item := range items {
		db.Create(&item)
	}
}
