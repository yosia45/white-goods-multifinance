package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserMoney struct {
	BaseModel
	UserID         uuid.UUID `json:"user_id" gorm:"not null;size:191"`
	Limit          float64   `json:"limit" gorm:"not null"`
	CurrentBalance float64   `json:"current_balance" gorm:"not null"`
}

func (um *UserMoney) BeforeCreate(tx *gorm.DB) (err error) {
	um.ID = uuid.New()
	um.CreatedAt = time.Now()
	um.UpdatedAt = time.Now()

	return
}
