package models

import (
	"github.com/google/uuid"
)

type Purchase struct {
	BaseModel
	UserLimitID    uuid.UUID     `json:"user_limit_id" gorm:"not null;size:191"`
	ItemTenorID    uuid.UUID     `json:"item_tenor_id" gorm:"not null;size:191"`
	MonthlyPayment float64       `json:"monthly_payment" gorm:"not null"`
	IsCompleted    bool          `json:"is_completed" gorm:"not null"`
	Transactions   []Transaction `json:"transactions" gorm:"foreignKey:PurchaseID"`
}
