package models

import "github.com/google/uuid"

type Purchase struct {
	BaseModel
	UserID         uuid.UUID     `json:"user_id" gorm:"not null;size:191"`
	TenorID        uuid.UUID     `json:"tenor_id" gorm:"not null;size:191"`
	ItemID         uuid.UUID     `json:"item_id" gorm:"not null;size:191"`
	MonthlyPayment float64       `json:"monthly_payment" gorm:"not null"`
	Status         string        `json:"status" gorm:"not null"` //ongoing or completed
	Transactions   []Transaction `json:"transactions" gorm:"foreignKey:PurchaseID"`
}
