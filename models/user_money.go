package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserMoney struct {
	BaseModel
	UserID       uuid.UUID `json:"user_id" gorm:"not null;size:191"`
	MonthlyLimit float64   `json:"monthly_limit" gorm:"not null;default:0"`
}

func (um *UserMoney) BeforeCreate(tx *gorm.DB) (err error) {
	um.ID = uuid.New()
	um.CreatedAt = time.Now()
	um.UpdatedAt = time.Now()

	return
}
