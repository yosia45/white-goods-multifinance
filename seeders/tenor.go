package seeders

import (
	"white-goods-multifinace/models"

	"gorm.io/gorm"
)

func SeedTenor(db *gorm.DB) {
	tenors := []models.Tenor{
		{
			Duration:  1,
			Interest:  0,
			IsDefault: true,
		},
		{
			Duration:  2,
			Interest:  0.01,
			IsDefault: true,
		},
		{
			Duration:  3,
			Interest:  0.02,
			IsDefault: true,
		},
		{
			Duration:  6,
			Interest:  0.03,
			IsDefault: true,
		},
	}

	for _, tenor := range tenors {
		db.Create(&tenor)
	}
}
