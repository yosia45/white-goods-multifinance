package seeders

import (
	"white-goods-multifinace/models"

	"gorm.io/gorm"
)

func SeedTenor(db *gorm.DB) {
	tenors := []models.Tenor{
		{
			Duration: 1,
		},
		{
			Duration: 2,
		},
		{
			Duration: 3,
		},
		{
			Duration: 6,
		},
	}

	for _, tenor := range tenors {
		db.Create(&tenor)
	}
}
