package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	BaseModel
	Name        string      `json:"name" gorm:"not null"`
	NormalPrice float64     `json:"normal_price" gorm:"not null"`
	ItemTenors  []ItemTenor `json:"item_tenors" gorm:"foreignKey:ItemID"`
}

func (i *Item) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = uuid.New()
	i.CreatedAt = time.Now()
	i.UpdatedAt = time.Now()

	return
}
