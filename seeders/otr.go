package seeders

import (
	"white-goods-multifinace/models"

	"gorm.io/gorm"
)

func SeedOTR(db *gorm.DB) {
	otrs := []models.OTR{
		{
			OTR: "white goods",
		},
		{
			OTR: "cars",
		},
		{
			OTR: "bikes",
		},
	}

	for _, otr := range otrs {
		db.Create(&otr)
	}
}
