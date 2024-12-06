package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemTenor struct {
	BaseModel
	ItemID   uuid.UUID `json:"item_id" gorm:"not null;size:191"`
	TenorID  uuid.UUID `json:"tenor_id" gorm:"not null;size:191"`
	Amount   float64   `json:"amount" gorm:"not null"`
	Interest float64   `json:"interest" gorm:"not null"`
}

func (it *ItemTenor) BeforeCreate(tx *gorm.DB) (err error) {
	it.ID = uuid.New()
	it.CreatedAt = time.Now()
	it.UpdatedAt = time.Now()

	return
}
