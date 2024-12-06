package seeders

import (
	"white-goods-multifinace/models"

	"gorm.io/gorm"
)

func SeedItems(db *gorm.DB) {
	items := []models.Item{
		{
			Name:        "Item 1",
			NormalPrice: 5000000,
			AdminFee:    100000,
			OTRID:       1,
		},
		{
			Name:        "Item 2",
			NormalPrice: 3000000,
			AdminFee:    80000,
			OTRID:       2,
		},
		{
			Name:        "Item 3",
			NormalPrice: 4000000,
			AdminFee:    90000,
			OTRID:       3,
		},
	}

	for _, item := range items {
		db.Create(&item)
	}
}
