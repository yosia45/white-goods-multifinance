package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	BaseModel
	PurchaseID    uuid.UUID `json:"purchase_id" gorm:"not null;size:191"`
	TotalAmount   float64   `json:"total_amount" gorm:"not null"`
	PaymentDate   time.Time `json:"payment_date" gorm:"not null"`
	InvoiceNumber string    `json:"invoice_number" gorm:"not null"`
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()

	return
}
