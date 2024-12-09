package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemTenor struct {
	BaseModel
	ItemID    uuid.UUID  `json:"item_id" gorm:"not null;size:191"`
	Item      Item       `json:"item" gorm:"foreignKey:ItemID"`
	TenorID   uuid.UUID  `json:"tenor_id" gorm:"not null;size:191"`
	Tenor     Tenor      `json:"tenor" gorm:"foreignKey:TenorID"`
	Interest  float64    `json:"interest" gorm:"not null"`
	Purchases []Purchase `json:"purchases" gorm:"foreignKey:ItemTenorID"`
}

func (it *ItemTenor) BeforeCreate(tx *gorm.DB) (err error) {
	it.ID = uuid.New()
	it.CreatedAt = time.Now()
	it.UpdatedAt = time.Now()

	return
}
