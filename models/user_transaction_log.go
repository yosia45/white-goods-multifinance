package models

import (
	"time"

	"github.com/google/uuid"
)

type UserTransactionLog struct {
	BaseModel
	UserPurchasingInformationID uuid.UUID `json:"user_purchasing_information_id" gorm:"not null;size:191"`
	InvoiceID                   string    `json:"invoice_id" gorm:"not null"`
	PaymentDate                 time.Time `json:"payment_date" gorm:"not null"`
	Status                      string    `json:"status" gorm:"not null"`
}
