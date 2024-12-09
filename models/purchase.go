package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Purchase struct {
	BaseModel
	UserLimitID uuid.UUID `json:"user_limit_id" gorm:"not null;size:191"`
	// UserLimit      UserLimit     `json:"user_limit" gorm:"foreignKey:UserLimitID"`
	ItemTenorID uuid.UUID `json:"item_tenor_id" gorm:"not null;size:191"`
	// ItemTenor      ItemTenor     `json:"item_tenor" gorm:"foreignKey:ItemTenorID"`
	MonthlyPayment float64       `json:"monthly_payment" gorm:"not null"`
	IsCompleted    bool          `json:"is_completed" gorm:"not null"`
	Transactions   []Transaction `json:"transactions" gorm:"foreignKey:PurchaseID"`
}

func (p *Purchase) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	return
}
