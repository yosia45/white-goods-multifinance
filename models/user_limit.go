package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserLimit struct {
	BaseModel
	UserID         uuid.UUID  `json:"user_id" gorm:"not null;size:191"`
	TenorID        uuid.UUID  `json:"tenor_id" gorm:"not null;size:191"`
	Limit          float64    `json:"limit" gorm:"not null;"`
	CurrentBalance float64    `json:"current_balance"`
	Purchases      []Purchase `json:"purchases" gorm:"foreignKey:UserLimitID"`
}

func (ul *UserLimit) BeforeCreate(tx *gorm.DB) (err error) {
	ul.ID = uuid.New()
	ul.CreatedAt = time.Now()
	ul.UpdatedAt = time.Now()

	return
}
