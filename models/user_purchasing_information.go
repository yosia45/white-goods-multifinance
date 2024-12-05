package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserPurchasingInformation struct {
	BaseModel
	UserID               uuid.UUID            `json:"user_id" gorm:"not null;size:191"`
	ItemTenorID          uuid.UUID            `json:"item_tenor_id" gorm:"not null;size:191"`
	Quantity             int                  `json:"quantity" gorm:"not null"`
	TotalNormalAmount    float64              `json:"total_normal_amount" gorm:"not null"`
	TotalRequiredAmount  float64              `json:"total_required_amount" gorm:"not null"`
	CurrentCreditPayment float64              `json:"current_credit_payment" gorm:"not null"`
	UserTransactionLogs  []UserTransactionLog `json:"user_transaction_logs" gorm:"foreignKey:UserPurchasingInformationID"`
}

func (upi *UserPurchasingInformation) BeforeCreate(tx *gorm.DB) (err error) {
	upi.ID = uuid.New()
	upi.CreatedAt = time.Now()
	upi.UpdatedAt = time.Now()

	return
}
