package models

import "gorm.io/gorm"

type OTR struct {
	gorm.Model
	OTR   string `json:"otr" gorm:"not null"`
	Items []Item `json:"items" gorm:"foreignKey:OTRID"`
}
