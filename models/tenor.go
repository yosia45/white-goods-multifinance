package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tenor struct {
	BaseModel
	Duration   int         `json:"duration" gorm:"not null"`
	Interest   float64     `json:"interest" gorm:"not null"`
	IsDefault  bool        `json:"is_default" gorm:"not null"`
	ItemTenors []ItemTenor `json:"item_tenors" gorm:"foreignKey:TenorID"`
}

func (t *Tenor) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()

	return
}
